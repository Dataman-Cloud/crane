package dockerclient

import (
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) InspectVolume(nodeId, name string) (*goclient.Volume, error) {
	return client.DockerClient(nodeId).InspectVolume(name)
}

func (client *RolexDockerClient) ListVolumes(nodeId string, opts goclient.ListVolumesOptions) ([]goclient.Volume, error) {
	return client.DockerClient(nodeId).ListVolumes(opts)
}

func (client *RolexDockerClient) CreateVolume(nodeId string, opts goclient.CreateVolumeOptions) (*goclient.Volume, error) {
	return client.DockerClient(nodeId).CreateVolume(opts)
}

func (client *RolexDockerClient) RemoveVolume(nodeId, name string) error {
	return client.DockerClient(nodeId).RemoveVolume(name)
}
