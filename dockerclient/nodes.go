package dockerclient

import (
	"encoding/json"
	"path"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
)

// NodeList returns the list of nodes.
func (client *RolexDockerClient) ListNode(opts types.NodeListOptions) ([]swarm.Node, error) {
	var nodes []swarm.Node

	content, err := client.HttpGet("nodes", nil, nil)
	if err != nil {
		return nodes, err
	}

	if err := json.Unmarshal(content, &nodes); err != nil {
		return nodes, err
	}

	return nodes, nil
}

// Inspect node returns the single node.
func (client *RolexDockerClient) InspectNode(nodeId string) (swarm.Node, error) {
	var node swarm.Node

	content, err := client.HttpGet(path.Join("nodes", nodeId), nil, nil)
	if err != nil {
		return node, err
	}

	if err := json.Unmarshal(content, &node); err != nil {
		return node, err
	}

	return node, nil
}

// Remove a single node
func (client *RolexDockerClient) RemoveNode(nodeId string) error {
	_, err := client.HttpDelete(path.Join("nodes", nodeId))
	if err != nil {
		return err
	}

	return nil
}

// docker info
func (client *RolexDockerClient) Info(nodeId string) (*goclient.DockerInfo, error) {
	return client.DockerClient(nodeId).Info()
}
