package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ConnectNetwork(id string, opts goclient.NetworkConnectionOptions) error {
	return client.DockerClient.ConnectNetwork(id, opts)
}

func (client *RolexDockerClient) CreateNetwork(opts goclient.CreateNetworkOptions) (*goclient.Network, error) {
	return client.DockerClient.CreateNetwork(opts)
}

func (client *RolexDockerClient) DisconnectNetwork(id string, opts goclient.NetworkConnectionOptions) error {
	return client.DockerClient.DisconnectNetwork(id, opts)
}

func (client *RolexDockerClient) InspectNetwork(id string) (*goclient.Network, error) {
	return client.DockerClient.NetworkInfo(id)
}

func (client *RolexDockerClient) ListNetworks(opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	return client.DockerClient.FilteredListNetworks(opts)
}

func (client *RolexDockerClient) RemoveNetwork(id string) error {
	return client.DockerClient.RemoveNetwork(id)
}
