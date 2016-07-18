package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

const (
	defaultNetworkDriver = "overlay"
)

func (client *RolexDockerClient) ConnectNetwork(id string, opts goclient.NetworkConnectionOptions) error {
	return client.SwarmClient().ConnectNetwork(id, opts)
}

func (client *RolexDockerClient) CreateNetwork(opts goclient.CreateNetworkOptions) (*goclient.Network, error) {
	return client.SwarmClient().CreateNetwork(opts)
}

func (client *RolexDockerClient) DisconnectNetwork(id string, opts goclient.NetworkConnectionOptions) error {
	return client.SwarmClient().DisconnectNetwork(id, opts)
}

func (client *RolexDockerClient) InspectNetwork(id string) (*goclient.Network, error) {
	return client.SwarmClient().NetworkInfo(id)
}

func (client *RolexDockerClient) ListNetworks(opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	return client.SwarmClient().FilteredListNetworks(opts)
}

func (client *RolexDockerClient) RemoveNetwork(id string) error {
	return client.SwarmClient().RemoveNetwork(id)
}

func (client *RolexDockerClient) ConnectNodeNetwork(nodeID, networkID string, opts goclient.NetworkConnectionOptions) error {
	return client.DockerClient(nodeID).ConnectNetwork(networkID, opts)
}

func (client *RolexDockerClient) DisconnectNodeNetwork(nodeID, networkID string, opts goclient.NetworkConnectionOptions) error {
	return client.DockerClient(nodeID).DisconnectNetwork(networkID, opts)
}

func (client *RolexDockerClient) InspectNodeNetwork(nodeID, networkID string) (*goclient.Network, error) {
	return client.DockerClient(nodeID).NetworkInfo(networkID)
}

func (client *RolexDockerClient) ListNodeNetworks(nodeID string, opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	return client.DockerClient(nodeID).FilteredListNetworks(opts)
}
