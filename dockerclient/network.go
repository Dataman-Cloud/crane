package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
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
