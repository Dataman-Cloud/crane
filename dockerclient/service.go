package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

type ServiceStatus struct {
	ID          string    `json:"ID"`
	Name        string    `json:"Name"`
	TaskRunning int       `json:"TaskRunning"`
	TaskTotal   int       `json:"TaskTotal"`
	Image       string    `json:"Images"`
	Command     string    `json:"Command"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

// Service scale instance
type ServiceScale struct {
	Scale uint64 `json:"Scale"`
}

const (
	TaskRunningState = "running"
)

// ServiceCreate creates a new Service.
func (client *RolexDockerClient) CreateService(service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	var response types.ServiceCreateResponse
	serviceParam, err := json.Marshal(service)
	if err != nil {
		return response, err
	}

	content, err := client.HttpPost(client.SwarmHttpEndpoint+"/services/create", nil, serviceParam, nil)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(content, &response); err != nil {
		return response, err
	}

	return response, nil
}

// ServiceList returns the list of services config
func (client *RolexDockerClient) ListServiceSpec(options types.ServiceListOptions) ([]swarm.Service, error) {
	var services []swarm.Service
	query := url.Values{}
	if options.Filter.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filter)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	content, err := client.HttpGet(client.SwarmHttpEndpoint+"/services", query, nil)
	if err != nil {
		return services, err
	}

	if err := json.Unmarshal(content, &services); err != nil {
		return services, err
	}

	return services, nil
}

// ListService return the list of service staus and core config
func (client *RolexDockerClient) ListService(options types.ServiceListOptions) ([]ServiceStatus, error) {
	services, err := client.ListServiceSpec(options)
	if err != nil {
		return nil, err
	}

	return client.GetServicesStatus(services)
}

// GetServicesStatus list services running status
func (client *RolexDockerClient) GetServicesStatus(services []swarm.Service) ([]ServiceStatus, error) {
	var servicesSt []ServiceStatus

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
		}
	}

	running := map[string]int{}
	for _, task := range tasks {
		if _, nodeActive := activeNodes[task.NodeID]; nodeActive && task.Status.State == TaskRunningState {
			running[task.ServiceID]++
		}
	}

	for _, service := range services {
		var taskTotal int
		if service.Spec.Mode.Replicated != nil && service.Spec.Mode.Replicated.Replicas != nil {
			taskTotal = int(*service.Spec.Mode.Replicated.Replicas)
		} else if service.Spec.Mode.Global != nil {
			taskTotal = len(activeNodes)
		}

		serviceSt := ServiceStatus{
			ID:          service.ID,
			Name:        service.Spec.Name,
			TaskRunning: running[service.ID],
			TaskTotal:   taskTotal,
			Image:       service.Spec.TaskTemplate.ContainerSpec.Image,
			Command:     strings.Join(service.Spec.TaskTemplate.ContainerSpec.Args, " "),
			CreatedAt:   service.CreatedAt,
			UpdatedAt:   service.UpdatedAt,
		}

		servicesSt = append(servicesSt, serviceSt)
	}

	return servicesSt, nil
}

// ServiceRemove kills and removes a service.
func (client *RolexDockerClient) RemoveService(serviceID string) error {
	_, err := client.HttpDelete(client.SwarmHttpEndpoint + "/services/" + serviceID)
	return err
}

// ServiceUpdate updates a Service.o
// TODO attention docker update
func (client *RolexDockerClient) UpdateService(serviceID string, version swarm.Version, service swarm.ServiceSpec, header map[string][]string) error {
	serviceParam, err := json.Marshal(service)
	if err != nil {
		return err
	}

	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	_, err = client.HttpPost(client.SwarmHttpEndpoint+"/services/"+serviceID+"/update", query, serviceParam, nil)
	if err != nil {
		return err
	}

	return nil
}

// ScaleService update service replicas
func (client *RolexDockerClient) ScaleService(serviceID string, serviceScale ServiceScale) error {
	service, err := client.InspectServiceWithRaw(serviceID)
	if err != nil {
		return err
	}

	serviceMode := &service.Spec.Mode
	if serviceMode.Replicated == nil {
		return fmt.Errorf("scale can only be used with replicated mode")
	}
	serviceMode.Replicated.Replicas = &serviceScale.Scale

	return client.UpdateService(service.ID, service.Version, service.Spec, nil)
}

// InspectServiceWithRaw returns the service information and the raw data.
func (client *RolexDockerClient) InspectServiceWithRaw(serviceID string) (swarm.Service, error) {
	var service swarm.Service

	content, err := client.HttpGet(client.SwarmHttpEndpoint+"/services/"+serviceID, nil, nil)
	if err != nil {
		return service, err
	}

	if err := json.Unmarshal(content, &service); err != nil {
		return service, err
	}

	return service, nil
}
