package dockerclient

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	distreference "github.com/docker/distribution/reference"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/docker/swarmkit/manager/scheduler"
	"github.com/docker/swarmkit/protobuf/ptypes"
)

var isValidName = regexp.MustCompile(`^[a-zA-Z0-9](?:[-_]*[A-Za-z0-9]+)*$`)

func validateResources(r *swarm.Resources) error {
	if r == nil {
		return nil
	}

	var errMsg string
	if r.NanoCPUs != 0 && r.NanoCPUs < 1e6 {
		errMsg = fmt.Sprintf("invalid cpu value %g: Must be at least %g", float64(r.NanoCPUs)/1e9, 1e6/1e9)
		return dmerror.NewError(CodeInvalidServiceNanoCPUs, errMsg)
	}

	if r.MemoryBytes != 0 && r.MemoryBytes < 4*1024*1024 {
		errMsg = fmt.Sprintf("invalid memory value %d: Must be at least 4MiB", r.MemoryBytes)
		return dmerror.NewError(CodeInvalidServiceMemoryBytes, errMsg)
	}
	return nil
}

func validateResourceRequirements(r *swarm.ResourceRequirements) error {
	if r == nil {
		return nil
	}
	if err := validateResources(r.Limits); err != nil {
		return err
	}
	if err := validateResources(r.Reservations); err != nil {
		return err
	}
	return nil
}

func validateRestartPolicy(rp *swarm.RestartPolicy) error {
	if rp == nil {
		return nil
	}

	var errMsg string
	if rp.Delay != nil {
		delay, err := ptypes.Duration(ptypes.DurationProto(*rp.Delay))
		if err != nil {
			return dmerror.NewError(CodeInvalidServiceDelay, err.Error())
		}
		if delay < 0 {
			errMsg = "TaskSpec: restart-delay cannot be negative"
			return dmerror.NewError(CodeInvalidServiceDelay, errMsg)
		}
	}

	if rp.Window != nil {
		win, err := ptypes.Duration(ptypes.DurationProto(*rp.Window))
		if err != nil {
			return dmerror.NewError(CodeInvalidServiceWindow, err.Error())
		}
		if win < 0 {
			errMsg = "TaskSpec: restart-window cannot be negative"
			return dmerror.NewError(CodeInvalidServiceWindow, errMsg)
		}
	}

	return nil
}

func validatePlacement(placement *swarm.Placement) error {
	if placement == nil {
		return nil
	}
	_, err := scheduler.ParseExprs(placement.Constraints)
	if err != nil {
		return dmerror.NewError(CodeInvalidServicePlacement, err.Error())
	}

	return nil
}

func validateUpdate(uc *swarm.UpdateConfig) error {
	if uc == nil {
		return nil
	}

	delay, err := ptypes.Duration(ptypes.DurationProto(uc.Delay))
	if err != nil {
		return dmerror.NewError(CodeInvalidServiceDelay, err.Error())
	}

	if delay < 0 {
		return dmerror.NewError(CodeInvalidServiceUpdateConfig, "TaskSpec: update-delay cannot be negative")
	}

	return nil
}

func validateTask(taskSpec swarm.TaskSpec) error {
	if err := validateResourceRequirements(taskSpec.Resources); err != nil {
		return err
	}

	if err := validateRestartPolicy(taskSpec.RestartPolicy); err != nil {
		return err
	}

	if err := validatePlacement(taskSpec.Placement); err != nil {
		return err
	}

	//TODO add this validate as soon
	//if taskSpec.GetRuntime() == nil {
	//	return grpc.Errorf(codes.InvalidArgument, "TaskSpec: missing runtime")
	//}

	//_, ok := taskSpec.GetRuntime().(*api.TaskSpec_Container)
	//if !ok {
	//	return grpc.Errorf(codes.Unimplemented, "RuntimeSpec: unimplemented runtime in service spec")
	//}

	//container := taskSpec.GetContainer()
	//if container == nil {
	//	return grpc.Errorf(codes.InvalidArgument, "ContainerSpec: missing in service spec")
	//}

	//if container.Image == "" {
	//	return grpc.Errorf(codes.InvalidArgument, "ContainerSpec: image reference must be provided")
	//}

	//if _, _, err := reference.Parse(container.Image); err != nil {
	//	return grpc.Errorf(codes.InvalidArgument, "ContainerSpec: %q is not a valid repository/tag", container.Image)
	//}
	return nil
}

func validateEndpointSpec(epSpec *swarm.EndpointSpec) error {
	// Endpoint spec is optional
	if epSpec == nil {
		return nil
	}

	if len(epSpec.Ports) > 0 && epSpec.Mode == swarm.ResolutionModeDNSRR {
		return dmerror.NewError(CodeInvalidServiceEndpoint, "EndpointSpec: ports can't be used with dnsrr mode")
	}

	portSet := make(map[swarm.PortConfig]struct{})
	for _, port := range epSpec.Ports {
		if _, ok := portSet[port]; ok {
			return dmerror.NewError(CodeInvalidServiceEndpoint, "EndpointSpec: duplicate ports provided")
		}

		portSet[port] = struct{}{}
	}

	return nil
}

func validateServiceSpec(spec *swarm.ServiceSpec) error {
	if spec == nil {
		return dmerror.NewError(CodeInvalidServiceSpec, "service spec must not null")
	}
	if err := validateAnnotations(spec.Annotations); err != nil {
		return err
	}
	if err := validateTask(spec.TaskTemplate); err != nil {
		return err
	}
	if err := validateUpdate(spec.UpdateConfig); err != nil {
		return err
	}
	if err := validateEndpointSpec(spec.EndpointSpec); err != nil {
		return err
	}

	if err := validateImageName(spec.TaskTemplate.ContainerSpec.Image); err != nil {
		return err
	}
	return nil
}

func validateAnnotations(m swarm.Annotations) error {
	if m.Name == "" {
		return dmerror.NewError(CodeInvalidServiceName, "meta: name must be provided")
	} else if !isValidName.MatchString(m.Name) {
		// if the name doesn't match the regex
		return dmerror.NewError(CodeInvalidServiceName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]")
	}
	return nil
}

func validateImageName(imageName string) error {
	_, err := distreference.ParseNamed(imageName)
	if err != nil {
		return dmerror.NewError(CodeInvalidImageName, err.Error())
	}
	return nil
}

// checkPortConflicts does a best effort to find if the passed in spec has port
// conflicts with existing services.
// `serviceID string` is the service ID of the spec in service update. If
// `serviceID` is not "", then conflicts check will be skipped against this
// check before create service `serviceId` is ""
func (client *RolexDockerClient) checkPortConflicts(reqPorts map[string]bool, serviceId string, existingServices []swarm.Service) error {
	for _, existingService := range existingServices {
		if serviceId != "" && serviceId == existingService.ID {
			continue
		}

		if existingService.Spec.EndpointSpec != nil {
			for _, pc := range existingService.Spec.EndpointSpec.Ports {
				portConflict := PortConflictToString(pc)
				if reqPorts[portConflict] {
					namespace := GetServicesNamespace(existingService.Spec)
					portConflictErr := &rolexerror.ServicePortConflictError{
						Name:          existingService.Spec.Name,
						Namespace:     namespace,
						PublishedPort: portConflict,
					}
					return &dmerror.DmError{Code: CodeGetServicePortConflictError, Err: portConflictErr}
				}
			}
		}

		for _, pc := range existingService.Endpoint.Ports {
			portConflict := PortConflictToString(pc)
			if reqPorts[portConflict] {
				namespace := GetServicesNamespace(existingService.Spec)
				portConflictErr := &rolexerror.ServicePortConflictError{
					Name:          existingService.Spec.Name,
					Namespace:     namespace,
					PublishedPort: portConflict,
				}
				return &dmerror.DmError{Code: CodeGetServicePortConflictError, Err: portConflictErr}
			}
		}

	}

	return nil
}

func PortConflictToString(pc swarm.PortConfig) string {
	port := strconv.FormatUint(uint64(pc.PublishedPort), 10)
	return port + "/" + string(pc.Protocol)
}

func (client *RolexDockerClient) CheckServicePortConflicts(spec *swarm.ServiceSpec, serviceId string) error {
	if spec.EndpointSpec == nil {
		return nil
	}

	reqPorts := make(map[string]bool)
	for _, pc := range spec.EndpointSpec.Ports {
		if pc.PublishedPort > 0 {
			reqPorts[PortConflictToString(pc)] = true
		}
	}

	if len(reqPorts) == 0 {
		return nil
	}

	existingServices, err := client.ListServiceSpec(types.ServiceListOptions{})
	if err != nil {
		return err
	}

	return client.checkPortConflicts(reqPorts, serviceId, existingServices)
}
