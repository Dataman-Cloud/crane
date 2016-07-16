package dockerclient

import (
	"github.com/Dataman-Cloud/rolex/util/rerror"
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListImages(nodeId string, opts goclient.ListImagesOptions) ([]goclient.APIImages, error) {
	images, err := client.DockerClient(nodeId).ListImages(opts)
	if err != nil {
		return images, rerror.NewRolexError(rerror.ENGINE_OPERATION_ERROR, err.Error())
	}

	return images, nil
}

func (client *RolexDockerClient) InspectImage(nodeId, name string) (*goclient.Image, error) {
	return client.DockerClient(nodeId).InspectImage(name)
}

func (client *RolexDockerClient) ImageHistory(nodeId, name string) ([]goclient.ImageHistory, error) {
	return client.DockerClient(nodeId).ImageHistory(name)
}
