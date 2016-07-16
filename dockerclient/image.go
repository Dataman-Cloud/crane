package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListImages(nodeID string, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	return client.DockerClient(nodeID).ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(nodeID, imageID string) (*goclient.Image, error) {
	return client.DockerClient(nodeID).InspectImage(imageID)
}

func (client *RolexDockerClient) ImageHistory(nodeID, imageID string) ([]goclient.ImageHistory, error) {
	return client.DockerClient(nodeID).ImageHistory(imageID)
}

// TODO add remoce image  option
func (client *RolexDockerClient) RemoveImage(nodeID, imageID string) error {
	return client.DockerClient(nodeID).RemoveImage(imageID)
}
