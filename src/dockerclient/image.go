package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) ListImages(ctx context.Context, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(ctx context.Context, imageID string) (*goclient.Image, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.InspectImage(imageID)
}

func (client *RolexDockerClient) ImageHistory(ctx context.Context, imageID string) ([]goclient.ImageHistory, error) {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return nil, err
	}
	return dockerClient.ImageHistory(imageID)
}

// TODO add remove image  option
func (client *RolexDockerClient) RemoveImage(ctx context.Context, imageID string) error {
	dockerClient, err := client.DockerClient(ctx)
	if err != nil {
		return err
	}
	return dockerClient.RemoveImage(imageID)
}
