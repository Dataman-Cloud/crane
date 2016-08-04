package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types/swarm"
)

// Inspect swarm cluster returns the swarm info
func (client *RolexDockerClient) InspectSwarm() (swarm.Swarm, error) {
	var swarmInfo swarm.Swarm

	content, err := client.HttpGet(client.SwarmHttpEndpoint+"/swarm", nil, nil)
	if err != nil {
		return swarmInfo, err
	}

	if err := json.Unmarshal(content, &swarmInfo); err != nil {
		return swarmInfo, err
	}

	return swarmInfo, nil
}
