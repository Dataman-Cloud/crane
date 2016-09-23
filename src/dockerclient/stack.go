package dockerclient

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

type Stacks []Stack

func (s Stacks) Len() int {
	return len(s)
}

func (s Stacks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Stacks) Less(i, j int) bool {
	return s[j].Services[0].CreatedAt.Unix() < s[i].Services[0].CreatedAt.Unix()
}

type Stack struct {
	// Name is the name of the stack
	Namespace string `json:"Namespace"`
	// Services is the number of the services
	ServiceCount int `json:"ServiceCount"`

	Services []ServiceStatus
}

// deploy a new stack
func (client *CraneDockerClient) DeployStack(bundle *model.Bundle) error {
	if bundle.Namespace == "" || !isValidName.MatchString(bundle.Namespace) {
		return cranerror.NewError(CodeInvalidStackName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]")
	}

	newNetworkMap, err := client.PretreatmentStack(*bundle)
	if err != nil {
		return err
	}

	return client.deployServices(bundle.Stack.Services, bundle.Namespace, newNetworkMap)
}

// before deploy stack we must verify all service spec params and check port conflict
// also we need check networks used by all of the servcie if the network is not existed
// created the network by the default param(network driver --overlay)
func (client *CraneDockerClient) PretreatmentStack(bundle model.Bundle) (map[string]bool, error) {
	// create network map and convert to slice for distinct network
	networkMap := make(map[string]bool)

	// all the publish port in the stack
	publishedPortMap := make(map[string]bool)

	for _, serviceSpec := range bundle.Stack.Services {
		if err := ValidateCraneServiceSpec(&serviceSpec); err != nil {
			return nil, err
		}

		for _, network := range serviceSpec.Networks {
			networkMap[network] = true
		}

		if serviceSpec.EndpointSpec != nil {
			for _, pc := range serviceSpec.EndpointSpec.Ports {
				if pc.PublishedPort > 0 {
					portConflictStr := PortConflictToString(pc)
					// have two service publish the same port
					if _, ok := publishedPortMap[portConflictStr]; ok {
						portConflictErr := &cranerror.ServicePortConflictError{
							Name:          serviceSpec.Name,
							Namespace:     bundle.Namespace,
							PublishedPort: portConflictStr,
						}
						return nil, &cranerror.CraneError{Code: CodeGetServicePortConflictError, Err: portConflictErr}
					}

					publishedPortMap[portConflictStr] = true
				}
			}
		}
	}

	// check stack need publish port is conflicted with exist services
	if len(publishedPortMap) > 0 {
		existingServices, err := client.ListServiceSpec(types.ServiceListOptions{})
		if err != nil {
			return nil, err
		}

		if err := checkPortConflicts(publishedPortMap, "", existingServices); err != nil {
			return nil, err
		}
	}

	// check if all network used by stack was exist, if not create it
	newNetworkMap, err := client.updateNetworks(networkMap, bundle.Namespace)
	if err != nil {
		return nil, err
	}

	return newNetworkMap, nil

}

// list all stack
func (client *CraneDockerClient) ListStack() (Stacks, error) {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace)
	services, err := client.ListServiceSpec(types.ServiceListOptions{Filter: filter})
	if err != nil {
		return nil, err
	}

	stackMap := make(map[string]*Stack, 0)
	for _, service := range services {
		labels := service.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok {
			log.Warnf("Cannot get label %s for service %s", labelNamespace, service.ID)
			continue
		}

		stack, ok := stackMap[name]
		if !ok {
			stackMap[name] = &Stack{
				Namespace:    name,
				ServiceCount: 1,
			}
		} else {
			stack.ServiceCount++
		}
	}

	var stacks Stacks
	for _, stack := range stackMap {
		stackServices, err := client.ListStackService(stack.Namespace, types.ServiceListOptions{})
		if err == nil {
			stack.Services = stackServices
		}
		stacks = append(stacks, *stack)
	}
	sort.Sort(stacks)

	return stacks, nil
}

// ListStackServices return list of service staus and core config in stack
func (client *CraneDockerClient) ListStackService(namespace string, opts types.ServiceListOptions) ([]ServiceStatus, error) {
	services, err := client.FilterServiceByStack(namespace, opts)
	if err != nil {
		return nil, err
	}

	return client.GetServicesStatus(services)
}

// Inspect stack get stack info
func (client *CraneDockerClient) InspectStack(namespace string) (*model.Bundle, error) {
	services, err := client.FilterServiceByStack(namespace, types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}

	stackServices := make(map[string]model.CraneServiceSpec)
	for _, swarmService := range services {
		stackServices[swarmService.Spec.Name] = client.ToCraneServiceSpec(swarmService.Spec)
	}

	return &model.Bundle{
		Namespace: namespace,
		Stack: model.BundleService{
			//TODO stack version is missing
			Services: stackServices,
		},
	}, nil
}

// remove all service and network in the stack
func (client *CraneDockerClient) RemoveStack(namespace string) error {
	services, err := client.FilterServiceByStack(namespace, types.ServiceListOptions{})
	if err != nil {
		return err
	}

	for _, service := range services {
		log.Info("begin to remove service ", service.Spec.Name)
		if err := client.RemoveService(service.ID); err != nil {
			return err
		}
	}

	networks, err := client.filterStackNetwork(namespace)
	if err != nil {
		return err
	}

	for _, network := range networks {
		log.Info("begin to remove network ", network.Name)
		if err := client.RemoveNetwork(network.ID); err != nil {
			return err
		}
	}

	if len(services) == 0 && len(networks) == 0 {
		return cranerror.NewError(CodeStackUnavailable, fmt.Sprintf("stack or network can't be empty %s ", namespace))
	}

	return nil
}

// filter service by stack name
func (client *CraneDockerClient) FilterServiceByStack(namespace string, opts types.ServiceListOptions) ([]swarm.Service, error) {
	if opts.Filter.Len() == 0 {
		opts.Filter = filters.NewArgs()
	}
	opts.Filter.Add("label", labelNamespace)
	services, err := client.ListServiceSpec(opts)
	if err != nil {
		return nil, err
	}

	var stackServices []swarm.Service
	for _, service := range services {
		labels := service.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok {
			log.Warnf("Cannot get label %s for service %s", labelNamespace, service.ID)
			continue
		}

		if name != namespace {
			continue
		}

		stackServices = append(stackServices, service)
	}

	return stackServices, nil
}

func (client *CraneDockerClient) GetStackGroup(namespace string) (uint64, error) {
	bundle, err := client.InspectStack(namespace)
	if err != nil {
		return 0, err
	}

	for _, service := range bundle.Stack.Services {
		for k, _ := range service.Labels {
			if strings.HasPrefix(k, "com.crane.permissions") {
				groupId, err := strconv.ParseUint(strings.Split(k, ".")[3], 10, 64)
				if err == nil {
					return groupId, nil
				}
			}
		}
	}
	return 0, errors.New("can't found stack groupid")
}

func (client *CraneDockerClient) updateNetworks(networks map[string]bool, namespace string) (map[string]bool, error) {
	existingNetworks, err := client.ListNetworks(docker.NetworkFilterOpts{})
	if err != nil {
		return nil, err
	}

	existingNetworkMap := make(map[string]docker.Network)
	for _, network := range existingNetworks {
		existingNetworkMap[network.Name] = network
	}

	createOpts := &docker.CreateNetworkOptions{
		Labels: client.getStackLabels(namespace, nil),
		Driver: defaultNetworkDriver,
		// docker TODO: remove when engine-api uses omitempty for IPAM
		IPAM: docker.IPAMOptions{Driver: "default"},
	}

	newNetworkMap := make(map[string]bool)
	for internalName := range networks {
		if _, exists := existingNetworkMap[internalName]; exists {
			newNetworkMap[internalName] = false
			continue
		}

		name := fmt.Sprintf("%s_%s", namespace, internalName)
		log.Infof("Creating network %s\n", name)
		createOpts.Name = name
		if _, err := client.CreateNetwork(*createOpts); err != nil {
			return newNetworkMap, err
		}
		newNetworkMap[internalName] = true
	}

	return newNetworkMap, nil
}

func convertNetworks(newNetworkMap map[string]bool, networks []string, namespace string, name string) []swarm.NetworkAttachmentConfig {
	nets := []swarm.NetworkAttachmentConfig{}
	for _, serviceNetwork := range networks {
		if isNew, ok := newNetworkMap[serviceNetwork]; ok && isNew {
			nets = append(nets, swarm.NetworkAttachmentConfig{
				Target:  namespace + "_" + serviceNetwork,
				Aliases: []string{name},
			})

		} else {
			nets = append(nets, swarm.NetworkAttachmentConfig{
				Target:  serviceNetwork,
				Aliases: []string{name},
			})
		}
	}
	return nets
}

func (client *CraneDockerClient) deployServices(services map[string]model.CraneServiceSpec, namespace string, newNetworkMap map[string]bool) error {
	existingServices, err := client.filterStackServices(namespace)
	if err != nil {
		return err
	}

	existingServiceMap := make(map[string]swarm.Service)
	for _, service := range existingServices {
		existingServiceMap[service.Spec.Name] = service
	}

	for internalName, service := range services {
		name := fmt.Sprintf("%s_%s", namespace, internalName)
		serviceSpec := swarm.ServiceSpec{
			Annotations: swarm.Annotations{
				Name:   name,
				Labels: client.getStackLabels(namespace, service.Labels),
			},
			Mode:         service.Mode,
			TaskTemplate: service.TaskTemplate,
			EndpointSpec: service.EndpointSpec,
			Networks:     convertNetworks(newNetworkMap, service.Networks, namespace, internalName),
			UpdateConfig: service.UpdateConfig,
		}

		createOpts := types.ServiceCreateOptions{}
		updateOpts := types.ServiceUpdateOptions{}
		if service.RegistryAuth != "" {
			registryAuth, err := EncodedRegistryAuth(service.RegistryAuth)
			if err != nil {
				return nil
			}
			createOpts.EncodedRegistryAuth = registryAuth
			updateOpts.EncodedRegistryAuth = registryAuth

			if serviceSpec.Labels == nil {
				serviceSpec.Labels = make(map[string]string)
			}

			serviceSpec.Annotations.Labels[LabelRegistryAuth] = service.RegistryAuth
		} else {
			// is safe to delete and not exist field
			delete(serviceSpec.Annotations.Labels, LabelRegistryAuth)
		}

		//TODO change service WorkingDir and User
		//cspec := &serviceSpec.TaskTemplate.ContainerSpec
		//if service.WorkingDir != nil {
		//	cspec.Dir = *service.WorkingDir
		//}

		//if service.User != nil {
		//	cspec.User = *service.User
		//}

		if service, exists := existingServiceMap[name]; exists {
			log.Infof("Updating service %s (id %s)", name, service.ID)

			if err := client.UpdateService(service.ID, service.Version, serviceSpec, updateOpts); err != nil {
				return err
			}
		} else {
			log.Infof("Creating service %s", name)
			if _, err := client.CreateService(serviceSpec, createOpts); err != nil {
				return err
			}
		}
	}

	return nil
}

// get stack labels
func (client *CraneDockerClient) getStackLabels(namespace string, labels map[string]string) map[string]string {
	if labels == nil {
		labels = make(map[string]string)
	}

	labels[labelNamespace] = namespace
	return labels
}

// split joint stack filter
func (client *CraneDockerClient) getStackFilter(namespace string) filters.Args {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace+"="+namespace)
	return filter
}

// get service by default stack labels
func (client *CraneDockerClient) filterStackServices(namespace string) ([]swarm.Service, error) {
	return client.ListServiceSpec(types.ServiceListOptions{Filter: client.getStackFilter(namespace)})
}

// get network by default filter
func (client *CraneDockerClient) filterStackNetwork(namespace string) ([]docker.Network, error) {
	filter := docker.NetworkFilterOpts{"label": map[string]bool{labelNamespace: true}}
	networks, err := client.ListNetworks(filter)
	if err != nil {
		return nil, err
	}

	var stackNetwork []docker.Network
	for _, network := range networks {
		if network.Labels == nil {
			continue
		}

		if name, ok := network.Labels[labelNamespace]; !ok || name != namespace {
			continue
		}

		stackNetwork = append(stackNetwork, network)
	}

	return stackNetwork, nil
}
