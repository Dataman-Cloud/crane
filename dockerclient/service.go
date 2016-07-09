package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
)

func (client *RolexDockerClient) ServiceCreate(service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
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
