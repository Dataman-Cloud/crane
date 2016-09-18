package dockerclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

var (
	ErrPermissionNotExists = errors.New("permission not exists")
)

type ServiceStatus struct {
	ID              string    `json:"ID"`
	Name            string    `json:"Name"`
	NumTasksRunning int       `json:"NumTasksRunning"`
	NumTasksTotal   int       `json:"NumTasksTotal"`
	Image           string    `json:"Image"`
	Command         string    `json:"Command"`
	CreatedAt       time.Time `json:"CreatedAt"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	LimitCpus       int64     `json:"LimitCpus"`
	LimitMems       int64     `json:"LimitMems"`
	ReserveCpus     int64     `json:"ReserveCpus"`
	ReserveMems     int64     `json:"ReserveMems"`
	IPs             []string  `json:"IPs"`
	Ports           []uint32  `json:"Ports"`
}

// scale a service request
type ServiceScale struct {
	NumTasks uint64 `json:"NumTasks"`
}

// ServiceCreate creates a new Service.
func (client *CraneDockerClient) CreateService(service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	var response types.ServiceCreateResponse
	var headers map[string][]string

	if options.EncodedRegistryAuth != "" {
		headers = map[string][]string{
			"X-Registry-Auth": []string{options.EncodedRegistryAuth},
		}
	}

	content, err := client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/services/create", nil, service, headers)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(content, &response); err != nil {
		return response, err
	}

	return response, nil
}

// ServiceList returns the list of services config
func (client *CraneDockerClient) ListServiceSpec(options types.ServiceListOptions) ([]swarm.Service, error) {
	var services []swarm.Service
	query := url.Values{}
	if options.Filter.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filter)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/services", query, nil)
	if err != nil {
		return services, err
	}

	if err := json.Unmarshal(content, &services); err != nil {
		return services, err
	}

	return services, nil
}

// ListService return the list of service staus and core config
func (client *CraneDockerClient) ListService(options types.ServiceListOptions) ([]ServiceStatus, error) {
	services, err := client.ListServiceSpec(options)
	if err != nil {
		return nil, err
	}

	return client.GetServicesStatus(services)
}

// GetServicesStatus list services running status
func (client *CraneDockerClient) GetServicesStatus(services []swarm.Service) ([]ServiceStatus, error) {
	var servicesSt []ServiceStatus
	var ips []string

	taskFilter := filters.NewArgs()
	for _, service := range services {
		taskFilter.Add("service", service.ID)
	}

	tasks, err := client.ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		return servicesSt, err
	}

	nodes, err := client.ListNode(types.NodeListOptions{})
	if err != nil {
		return servicesSt, err
	}

	activeNodes := make(map[string]struct{})
	for _, node := range nodes {
		if node.Status.State == swarm.NodeStateReady {
			activeNodes[node.ID] = struct{}{}
			if endpoint, ok := node.Spec.Labels["crane.reserved.node.endpoint"]; ok {
				ips = append(ips, endpoint)
			}
		}
	}

	runningTasks := map[string]int{}
	for _, task := range tasks {
		if _, nodeActive := activeNodes[task.NodeID]; nodeActive && task.Status.State == TaskRunningState {
			runningTasks[task.ServiceID]++
		}
	}

	for _, service := range services {
		var taskTotal int
		if service.Spec.Mode.Replicated != nil && service.Spec.Mode.Replicated.Replicas != nil {
			taskTotal = int(*service.Spec.Mode.Replicated.Replicas)
		} else if service.Spec.Mode.Global != nil {
			taskTotal = len(activeNodes)
		}

		var limitCpus int64
		var limitMems int64
		var reserveCpus int64
		var reserveMems int64
		if service.Spec.TaskTemplate.Resources != nil {
			limitCpus = int64(runningTasks[service.ID]) * service.Spec.TaskTemplate.Resources.Limits.NanoCPUs
			limitMems = int64(runningTasks[service.ID]) * service.Spec.TaskTemplate.Resources.Limits.MemoryBytes
			reserveCpus = int64(taskTotal) * service.Spec.TaskTemplate.Resources.Reservations.NanoCPUs
			reserveMems = int64(taskTotal) * service.Spec.TaskTemplate.Resources.Reservations.MemoryBytes
		}

		var ports []uint32
		if service.Spec.EndpointSpec != nil {
			for _, port := range service.Spec.EndpointSpec.Ports {
				ports = append(ports, port.PublishedPort)
			}
		}

		serviceSt := ServiceStatus{
			ID:              service.ID,
			Name:            service.Spec.Name,
			NumTasksRunning: runningTasks[service.ID],
			NumTasksTotal:   taskTotal,
			Image:           service.Spec.TaskTemplate.ContainerSpec.Image,
			Command:         strings.Join(service.Spec.TaskTemplate.ContainerSpec.Args, " "),
			CreatedAt:       service.CreatedAt,
			UpdatedAt:       service.UpdatedAt,
			LimitCpus:       limitCpus,
			LimitMems:       limitMems,
			ReserveCpus:     reserveCpus,
			ReserveMems:     reserveMems,
			IPs:             ips,
			Ports:           ports,
		}

		servicesSt = append(servicesSt, serviceSt)
	}

	return servicesSt, nil
}

// ServiceRemove kills and removes a service.
func (client *CraneDockerClient) RemoveService(serviceID string) error {
	_, err := client.sharedHttpClient.DELETE(nil, client.swarmManagerHttpEndpoint+"/services/"+serviceID, nil, nil)
	return err
}

// ServiceUpdate updates a Service.o
func (client *CraneDockerClient) UpdateService(serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) error {
	var headers map[string][]string
	if options.EncodedRegistryAuth != "" {
		headers = map[string][]string{
			"X-Registry-Auth": []string{options.EncodedRegistryAuth},
		}
	}

	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	if _, err := client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/services/"+serviceID+"/update", query, service, headers); err != nil {
		return err
	}

	return nil
}

func (client *CraneDockerClient) UpdateServiceAutoOption(serviceID string, version swarm.Version, service swarm.ServiceSpec) error {
	updateOpts := types.ServiceUpdateOptions{}
	if service.Annotations.Labels != nil {
		if registryAuth, ok := service.Annotations.Labels[LabelRegistryAuth]; ok {
			encodeRegistryAuth, err := EncodedRegistryAuth(registryAuth)
			if err != nil {
				return nil
			}
			updateOpts.EncodedRegistryAuth = encodeRegistryAuth
		}
	}

	return client.UpdateService(serviceID, version, service, updateOpts)

}

// ScaleService update service replicas
func (client *CraneDockerClient) ScaleService(serviceID string, serviceScale ServiceScale) error {
	service, err := client.InspectServiceWithRaw(serviceID)
	if err != nil {
		return err
	}

	serviceMode := &service.Spec.Mode
	if serviceMode.Replicated == nil {
		return fmt.Errorf("scale can only be used with replicated mode")
	}
	serviceMode.Replicated.Replicas = &serviceScale.NumTasks

	return client.UpdateServiceAutoOption(service.ID, service.Version, service.Spec)
}

// InspectServiceWithRaw returns the service information and the raw data.
func (client *CraneDockerClient) InspectServiceWithRaw(serviceID string) (swarm.Service, error) {
	var service swarm.Service

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/services/"+serviceID, nil, nil)
	if err != nil {
		return service, err
	}

	if err := json.Unmarshal(content, &service); err != nil {
		return service, err
	}

	return service, nil
}

// grant service permissions
func (client *CraneDockerClient) ServiceAddLabel(serviceID string, labels map[string]string) error {
	service, err := client.InspectServiceWithRaw(serviceID)
	if err != nil {
		return err
	}

	for k, v := range labels {
		service.Spec.Labels[k] = v
	}

	return client.UpdateServiceAutoOption(service.ID, service.Version, service.Spec)
}

// revoke service permissions
func (client *CraneDockerClient) ServiceRemoveLabel(serviceID string, labels []string) error {
	service, err := client.InspectServiceWithRaw(serviceID)
	if err != nil {
		return err
	}

	for _, label := range labels {
		delete(service.Spec.Labels, label)
	}

	return client.UpdateServiceAutoOption(service.ID, service.Version, service.Spec)
}

func (client *CraneDockerClient) GetServiceNetworkNames(networkAttachmentConfigs []swarm.NetworkAttachmentConfig) []string {
	networkNameList := []string{}
	for _, networkAttachmentConfig := range networkAttachmentConfigs {
		networkInfo, err := client.InspectNetwork(networkAttachmentConfig.Target)
		if err != nil {
			log.Warnf("convert service network got error: %f", err.Error())
			continue
		}

		networkNameList = append(networkNameList, networkInfo.Name)
	}

	return networkNameList
}
