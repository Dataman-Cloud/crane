package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListImages(opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	return client.DockerClient("placeholder").ListImages(opts)
}

func (client *RolexDockerClient) InspectImage(name string) (*goclient.Image, error) {
	return client.DockerClient("placeholder").InspectImage(name)
}

func (client *RolexDockerClient) ImageHistory(name string) ([]goclient.ImageHistory, error) {
	return client.DockerClient("placeholder").ImageHistory(name)
}
