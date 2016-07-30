package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) ListImages(ctx context.Context, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	return client.DockerClient(ctx).ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(ctx context.Context, imageID string) (*goclient.Image, error) {
	return client.DockerClient(ctx).InspectImage(imageID)
}

func (client *RolexDockerClient) ImageHistory(ctx context.Context, imageID string) ([]goclient.ImageHistory, error) {
	return client.DockerClient(ctx).ImageHistory(imageID)
}

// TODO add remoce image  option
func (client *RolexDockerClient) RemoveImage(ctx context.Context, imageID string) error {
	return client.DockerClient(ctx).RemoveImage(imageID)
}
