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
	return client.SwarmManager().ConnectNetwork(id, opts)
}

func (client *RolexDockerClient) CreateNetwork(opts goclient.CreateNetworkOptions) (*goclient.Network, error) {
	return client.SwarmManager().CreateNetwork(opts)
}

func (client *RolexDockerClient) DisconnectNetwork(id string, opts goclient.NetworkConnectionOptions) error {
	return client.SwarmManager().DisconnectNetwork(id, opts)
}

func (client *RolexDockerClient) InspectNetwork(id string) (*goclient.Network, error) {
	return client.SwarmManager().NetworkInfo(id)
}

func (client *RolexDockerClient) ListNetworks(opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	return client.SwarmManager().FilteredListNetworks(opts)
}

func (client *RolexDockerClient) RemoveNetwork(id string) error {
	err := client.SwarmManager().RemoveNetwork(id)
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
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}

	err = swarmNode.ConnectNetwork(networkID, opts)
	if err != nil {
		err = ToRolexError(err)
	}

	return err
}

func (client *RolexDockerClient) DisconnectNodeNetwork(ctx context.Context, networkID string, opts goclient.NetworkConnectionOptions) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}

	err = swarmNode.DisconnectNetwork(networkID, opts)
	if err != nil {
		err = ToRolexError(err)
	}

	return err
}

func (client *RolexDockerClient) InspectNodeNetwork(ctx context.Context, networkID string) (*goclient.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}

	network, err := swarmNode.NetworkInfo(networkID)
	if err != nil {
		err = ToRolexError(err)
	}

	return network, err
}

func (client *RolexDockerClient) ListNodeNetworks(ctx context.Context, opts goclient.NetworkFilterOpts) ([]goclient.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.FilteredListNetworks(opts)
}

func (client *RolexDockerClient) CreateNodeNetwork(ctx context.Context, opts goclient.CreateNetworkOptions) (*goclient.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.CreateNetwork(opts)
}
