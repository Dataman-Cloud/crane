package dockerclient

import (
	"strings"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *CraneDockerClient) ConnectNetwork(id string, opts docker.NetworkConnectionOptions) error {
	return client.SwarmManager().ConnectNetwork(id, opts)
}

func (client *CraneDockerClient) CreateNetwork(opts docker.CreateNetworkOptions) (*docker.Network, error) {
	if opts.Name == "" || !isValidName.MatchString(opts.Name) {
		return nil, cranerror.NewError(CodeInvalidNetworkName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9] are allowed")
	}

	return client.SwarmManager().CreateNetwork(opts)
}

func (client *CraneDockerClient) DisconnectNetwork(id string, opts docker.NetworkConnectionOptions) error {
	return client.SwarmManager().DisconnectNetwork(id, opts)
}

func (client *CraneDockerClient) InspectNetwork(id string) (*docker.Network, error) {
	return client.SwarmManager().NetworkInfo(id)
}

func (client *CraneDockerClient) ListNetworks(opts docker.NetworkFilterOpts) ([]docker.Network, error) {
	return client.SwarmManager().FilteredListNetworks(opts)
}

func (client *CraneDockerClient) RemoveNetwork(id string) error {
	err := client.SwarmManager().RemoveNetwork(id)
	if err != nil {
		if strings.Contains(err.Error(), "API error (403)") {
			return cranerror.NewError(CodeNetworkPredefined, err.Error())
		} else {
			return err
		}
	}

	return nil
}

func (client *CraneDockerClient) ConnectNodeNetwork(ctx context.Context, networkID string, opts docker.NetworkConnectionOptions) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}

	err = swarmNode.ConnectNetwork(networkID, opts)
	if err != nil {
		err = ToCraneError(err)
	}

	return err
}

func (client *CraneDockerClient) DisconnectNodeNetwork(ctx context.Context, networkID string, opts docker.NetworkConnectionOptions) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}

	err = swarmNode.DisconnectNetwork(networkID, opts)
	if err != nil {
		err = ToCraneError(err)
	}

	return err
}

func (client *CraneDockerClient) InspectNodeNetwork(ctx context.Context, networkID string) (*docker.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}

	network, err := swarmNode.NetworkInfo(networkID)
	if err != nil {
		err = ToCraneError(err)
	}

	return network, err
}

func (client *CraneDockerClient) ListNodeNetworks(ctx context.Context, opts docker.NetworkFilterOpts) ([]docker.Network, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.FilteredListNetworks(opts)
}

func (client *CraneDockerClient) CreateNodeNetwork(ctx context.Context, opts docker.CreateNetworkOptions) (*docker.Network, error) {
	if opts.Name == "" || !isValidName.MatchString(opts.Name) {
		return nil, cranerror.NewError(CodeInvalidNetworkName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9] are allowed")
	}

	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.CreateNetwork(opts)
}
