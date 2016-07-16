package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListImages(nodeId string, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	return client.DockerClient(nodeId).ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(nodeId, imageId string) (*goclient.Image, error) {
	return client.DockerClient(nodeId).InspectImage(imageId)
}

func (client *RolexDockerClient) ImageHistory(nodeId, imageId string) ([]goclient.ImageHistory, error) {
	return client.DockerClient(nodeId).ImageHistory(imageId)
}
