package dockerclient

import (
	"strings"

	"github.com/Dataman-Cloud/crane/src/utils/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) ConnectNetwork(id string, opts docker.NetworkConnectionOptions) error {
	return client.SwarmManager().ConnectNetwork(id, opts)
}

func (client *RolexDockerClient) CreateNetwork(opts docker.CreateNetworkOptions) (*docker.Network, error) {
	if opts.Name == "" || !isValidName.MatchString(opts.Name) {
		return nil, rolexerror.NewError(CodeInvalidNetworkName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9] are allowed")
	}

	return client.SwarmManager().CreateNetwork(opts)
}

func (client *RolexDockerClient) DisconnectNetwork(id string, opts docker.NetworkConnectionOptions) error {
	return client.SwarmManager().DisconnectNetwork(id, opts)
}

func (client *RolexDockerClient) InspectNetwork(id string) (*docker.Network, error) {
	return client.SwarmManager().NetworkInfo(id)
}

func (client *RolexDockerClient) ListNetworks(opts docker.NetworkFilterOpts) ([]docker.Network, error) {
	return client.SwarmManager().FilteredListNetworks(opts)
}

func (client *RolexDockerClient) RemoveNetwork(id string) error {
	err := client.SwarmManager().RemoveNetwork(id)
	if err != nil {
		if strings.Contains(err.Error(), "API error (403)") {
			return rolexerror.NewError(CodeNetworkPredefined, err.Error())
		} else {
			return err
		}
	}

	return nil
}

func (client *RolexDockerClient) ConnectNodeNetwork(ctx context.Context, networkID string, opts docker.NetworkConnectionOptions) error {
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

func (client *RolexDockerClient) DisconnectNodeNetwork(ctx context.Context, networkID string, opts docker.NetworkConnectionOptions) error {
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

func (client *RolexDockerClient) InspectNodeNetwork(ctx context.Context, networkID string) (*docker.Network, error) {
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

func (client *RolexDockerClient) ListNodeNetworks(ctx context.Context, opts docker.NetworkFilterOpts) ([]docker.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.FilteredListNetworks(opts)
}

func (client *RolexDockerClient) CreateNodeNetwork(ctx context.Context, opts docker.CreateNetworkOptions) (*docker.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.CreateNetwork(opts)
}
