package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
)

// Inspect swarm cluster returns the swarm info
func (client *RolexDockerClient) InspectSwarm() (swarm.Swarm, error) {
	var swarmInfo swarm.Swarm

	content, err := client.HttpGet(client.swarmManagerHttpEndpoint+"/swarm", nil, nil)
	if err != nil {
		return swarmInfo, err
	}

	if err := json.Unmarshal(content, &swarmInfo); err != nil {
		return swarmInfo, err
	}

	return swarmInfo, nil
}

// ping to test swarmManager connection
func (client *RolexDockerClient) Ping() error {
	return client.swarmManager.Ping()
}

// Get Manager information, equal to client cmd `docker info` on the manager node
func (client *RolexDockerClient) ManagerInfo() (types.Info, error) {
	var systemInfo types.Info

	content, err := client.HttpGet(client.swarmManagerHttpEndpoint+"/info", nil, nil)
	if err != nil {
		return systemInfo, err
	}

	if err := json.Unmarshal(content, &systemInfo); err != nil {
		return systemInfo, err
	}

	return systemInfo, nil
}
