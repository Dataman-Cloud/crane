package dockerclient

import (
	"strings"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
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
	err := client.SwarmClient().RemoveNetwork(id)
	if err != nil {
		if strings.Contains(err.Error(), "API error (403)") {
			return rolexerror.NewRolexError(rolexerror.CodeNetworkPredefined, err.Error())
		} else {
			return err
		}
	}

	return nil
}

func (client *RolexDockerClient) ConnectNodeNetwork(ctx context.Context, networkID string, opts goclient.NetworkConnectionOptions) error {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return err
	}
	return dockerClient.ConnectNetwork(networkID, opts)
}

func (client *RolexDockerClient) DisconnectNodeNetwork(ctx context.Context, networkID string, opts goclient.NetworkConnectionOptions) error {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return err
	}
	return dockerClient.DisconnectNetwork(networkID, opts)
}

func (client *RolexDockerClient) InspectNodeNetwork(ctx context.Context, networkID string) (*goclient.Network, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.NetworkInfo(networkID)
}

func (client *RolexDockerClient) ListNodeNetworks(ctx context.Context, opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.FilteredListNetworks(opts)
}

func (client *RolexDockerClient) CreateNodeNetwork(ctx context.Context, opts goclient.CreateNetworkOptions) (*goclient.Network, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.CreateNetwork(opts)
}
