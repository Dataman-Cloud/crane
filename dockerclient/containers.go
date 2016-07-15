package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListContainers(opts goclient.ListContainersOptions) ([]goclient.APIContainers, error) {
	return client.DockerClient("placeholder").ListContainers(opts)
}

func (client *RolexDockerClient) InspectContainer(id string) (*goclient.Container, error) {
	return client.DockerClient("placeholder").InspectContainer(id)
}

func (client *RolexDockerClient) RemoveContainer(opts goclient.RemoveContainerOptions) error {
	return client.DockerClient("placeholder").RemoveContainer(opts)
}

func (client *RolexDockerClient) KillContainer(opts goclient.KillContainerOptions) error {
	return client.DockerClient("placeholder").KillContainer(opts)
}

func (client *RolexDockerClient) DiffContainer(containerID string) ([]types.ContainerChange, error) {
	var changes []types.ContainerChange

	content, err := client.HttpGet("/containers/" + containerID + "/changes")
	if err != nil {
		return changes, err
	}

	if err := json.Unmarshal(content, &changes); err != nil {
		return changes, err
	}

	return changes, nil
}
