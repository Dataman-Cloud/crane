package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) InspectVolume(ctx context.Context, name string) (*goclient.Volume, error) {
	return client.DockerClient(ctx).InspectVolume(name)
}

func (client *RolexDockerClient) ListVolumes(ctx context.Context, opts goclient.ListVolumesOptions) ([]goclient.Volume, error) {
	return client.DockerClient(ctx).ListVolumes(opts)
}

func (client *RolexDockerClient) CreateVolume(ctx context.Context, opts goclient.CreateVolumeOptions) (*goclient.Volume, error) {
	return client.DockerClient(ctx).CreateVolume(opts)
}

func (client *RolexDockerClient) RemoveVolume(ctx context.Context, name string) error {
	return client.DockerClient(ctx).RemoveVolume(name)
}
