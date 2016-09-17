package dockerclient

import (
	"encoding/json"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
)

// Inspect swarm cluster returns the swarm info
func (client *CraneDockerClient) InspectSwarm() (swarm.Swarm, error) {
	var swarmInfo swarm.Swarm

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/swarm", nil, nil)
	if err != nil {
		return swarmInfo, err
	}

	if err := json.Unmarshal(content, &swarmInfo); err != nil {
		return swarmInfo, err
	}

	return swarmInfo, nil
}

// ping to test swarmManager connection
func (client *CraneDockerClient) Ping() error {
	return client.swarmManager.Ping()
}

// Get Manager information, equal to client cmd `docker info` on the manager node
func (client *CraneDockerClient) ManagerInfo() (types.Info, error) {
	var systemInfo types.Info

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/info", nil, nil)
	if err != nil {
		return systemInfo, cranerror.NewError(CodeGetManagerInfoError, err.Error())
	}

	if err := json.Unmarshal(content, &systemInfo); err != nil {
		return systemInfo, cranerror.NewError(CodeGetManagerInfoError, err.Error())
	}

	return systemInfo, nil
}
