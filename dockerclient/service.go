package dockerclient

import (
	"encoding/json"
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

	content, err := client.HttpPost("services/create", serviceParam)
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
	content, err := client.HttpGet("/services")
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
		serviceSt := ServiceStatus{
			ID:          service.ID,
			Name:        service.Spec.Name,
			TaskRunning: running[service.ID],
			TaskTotal:   len(activeNodes),
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
	_, err := client.HttpDelete("/services" + serviceID)
	return err
}

// ServiceUpdate updates a Service.o
// TODO attention docker update
func (client *RolexDockerClient) UpdateService(serviceID string, version swarm.Version, service swarm.ServiceSpec, header map[string][]string) error {
	serviceParam, err := json.Marshal(service)
	if err != nil {
		return err
	}

	_, err = client.HttpPost("services/"+serviceID+"/update", serviceParam)
	if err != nil {
		return err
	}

	return nil
}
