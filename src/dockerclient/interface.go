package dockerclient

import (
	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	opt "github.com/Dataman-Cloud/rolex/src/model"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

type DockerClientInterface interface {
	Ping() error

	NodeList(opts types.NodeListOptions) ([]swarm.Node, error)
	NodeInspect(nodeId string) (swarm.Node, error)
	NodeRemove(nodeId string) error
	UpdateNode(nodeId string, opts opt.UpdateOptions) error

	InspectContainer(id string) (*goclient.Container, error)
	ListContainers(opts goclient.ListContainersOptions) ([]goclient.APIContainers, error)
	RemoveContainer(opts goclient.RemoveContainerOptions) error
	LogsContainer(nodeId, containerId string, message chan string)
	StatsContainer(nodeId, containerId string, stats chan *model.ContainerStat)

	ConnectNetwork(id string, opts goclient.NetworkConnectionOptions) error
	CreateNetwork(opts goclient.CreateNetworkOptions) (*goclient.Network, error)
	DisconnectNetwork(id string, opts goclient.NetworkConnectionOptions) error
	InspectNetwork(id string) (*goclient.Network, error)
	ListNetworks(opts goclient.NetworkFilterOpts) ([]goclient.Network, error)
	RemoveNetwork(id string) error

	CreateNodeNetwork(ctx context.Context, opts goclient.CreateNetworkOptions) (*goclient.Network, error)
	ConnectNodeNetwork(ctx context.Context, networkID string, opts goclient.NetworkConnectionOptions) error
	DisconnectNodeNetwork(ctx context.Context, networkID string, opts goclient.NetworkConnectionOptions) error
	InspectNodeNetwork(ctx context.Context, networkID string) (*goclient.Network, error)
	ListNodeNetworks(ctx context.Context, opts goclient.NetworkFilterOpts) ([]goclient.Network, error)

	InspectVolume(nodeId, name string) (*goclient.Volume, error)
	ListVolumes(nodeId string, opts goclient.ListVolumesOptions) ([]goclient.Volume, error)
	CreateVolume(nodeId string, opts goclient.CreateVolumeOptions) (*goclient.Volume, error)
	RemoveVolume(nodeId string, name string) error

	ListImages(nodeId string, opts goclient.ListImagesOptions) ([]goclient.APIImages, error)
	InspectImage(nodeId, imageId string) (*goclient.Image, error)
	ImageHistory(nodeId, imageId string) ([]goclient.ImageHistory, error)

	DeployStack(bundle *model.Bundle) error
	ListStack() ([]Stack, error)
	ListStackService(namespace string, opts types.ServiceListOptions) ([]ServiceStatus, error)
	InspectStack(namespace string) (*model.Bundle, error)
	RemoveStack(namespace string) error
	FilterServiceByStack(namespace string, opts types.ServiceListOptions) ([]swarm.Service, error)
	ConvertStackService(swarmService swarm.ServiceSpec) model.RolexService

	CreateService(service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error)
	ListServiceSpec(options types.ServiceListOptions) ([]swarm.Service, error)
	ListService(options types.ServiceListOptions) ([]ServiceStatus, error)
	GetServicesStatus(services []swarm.Service) ([]ServiceStatus, error)
	RemoveService(serviceID string) error
	UpdateService(serviceID string, version swarm.Version, service swarm.ServiceSpec, header map[string][]string) error
	ScaleService(serviceID string, serviceScale ServiceScale) error
	InspectServiceWithRaw(serviceID string) (swarm.Service, error)
	ServiceAddLabel(serviceID string, labels map[string]string) error
	ServiceRemoveLabel(serviceID string, labels []string) error
}
