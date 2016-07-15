package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListImages(nodeId string, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	return client.DockerClient(nodeId).ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(nodeId, name string) (*goclient.Image, error) {
	return client.DockerClient(nodeId).InspectImage(name)
}

func (client *RolexDockerClient) ImageHistory(nodeId, name string) ([]goclient.ImageHistory, error) {
	return client.DockerClient(nodeId).ImageHistory(name)
}
