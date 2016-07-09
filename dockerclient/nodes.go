package dockerclient

import (
	"encoding/json"
	"path"

	"github.com/docker/engine-api/types/swarm"
)

// NodeList returns the list of nodes.
func (client *RolexDockerClient) NodeList(opts NodeListOptions) ([]swarm.Node, error) {
	var nodes []swarm.Node

	content, err := client.HttpGet("nodes")
	if err != nil {
		return nodes, err
	}

	if err := json.Unmarshal(content, &nodes); err != nil {
		return nodes, err
	}

	return nodes, nil
}

// Inspect node returns the single node.
func (client *RolexDockerClient) NodeInspect(nodeId string) (swarm.Node, error) {
	var node swarm.Node

	content, err := client.HttpGet(path.Join("nodes", nodeId))
	if err != nil {
		return node, err
	}

	if err := json.Unmarshal(content, &node); err != nil {
		return node, err
	}

	return node, nil
}

// Remove a single node
func (client *RolexDockerClient) NodeRemove(nodeId string) error {
	content, err := client.HttpDelete(path.Join("nodes", nodeId))
	if err != nil {
		return err
	}

	return nil
}
