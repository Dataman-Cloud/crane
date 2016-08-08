package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) ListImages(ctx context.Context, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(ctx context.Context, imageID string) (*goclient.Image, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.InspectImage(imageID)
}

func (client *RolexDockerClient) ImageHistory(ctx context.Context, imageID string) ([]goclient.ImageHistory, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.ImageHistory(imageID)
}

// TODO add remove image  option
func (client *RolexDockerClient) RemoveImage(ctx context.Context, imageID string) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}
	return swarmNode.RemoveImage(imageID)
}
