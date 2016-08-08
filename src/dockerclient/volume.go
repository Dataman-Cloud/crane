package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) InspectVolume(ctx context.Context, name string) (*goclient.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.InspectVolume(name)
}

func (client *RolexDockerClient) ListVolumes(ctx context.Context, opts goclient.ListVolumesOptions) ([]goclient.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.ListVolumes(opts)
}

func (client *RolexDockerClient) CreateVolume(ctx context.Context, opts goclient.CreateVolumeOptions) (*goclient.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.CreateVolume(opts)
}

func (client *RolexDockerClient) RemoveVolume(ctx context.Context, name string) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}
	return swarmNode.RemoveVolume(name)
}
