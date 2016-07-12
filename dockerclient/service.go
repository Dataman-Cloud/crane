package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
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

// ServiceList returns the list of services.
func (client *RolexDockerClient) ListService(options types.ServiceListOptions) ([]swarm.Service, error) {
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
