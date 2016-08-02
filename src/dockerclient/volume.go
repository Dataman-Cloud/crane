package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) InspectVolume(ctx context.Context, name string) (*goclient.Volume, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.InspectVolume(name)
}

func (client *RolexDockerClient) ListVolumes(ctx context.Context, opts goclient.ListVolumesOptions) ([]goclient.Volume, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.ListVolumes(opts)
}

func (client *RolexDockerClient) CreateVolume(ctx context.Context, opts goclient.CreateVolumeOptions) (*goclient.Volume, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.CreateVolume(opts)
}

func (client *RolexDockerClient) RemoveVolume(ctx context.Context, name string) error {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return err
	}
	return dockerClient.RemoveVolume(name)
}
