package dockerclient

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
)

type DockerClientInterface interface {
	Ping() error

	NodeList(opts types.NodeListOptions) ([]swarm.Node, error)
	NodeInspect(nodeId string) (swarm.Node, error)
	NodeRemove(nodeId string) error

	InspectContainer(id string) (*goclient.Container, error)
	ListContainers(opts goclient.ListContainersOptions) ([]goclient.APIContainers, error)
	RemoveContainer(opts goclient.RemoveContainerOptions) error
}
