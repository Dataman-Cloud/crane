package dockerclient

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/client/bundlefile"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
)

const (
	defaultNetworkDriver = "overlay"
	labelNamespace       = "com.docker.stack.namespace"
)

// bundle stores the contents of services and stack name
type Bundle struct {
	bundlefile.Bundlefile
	NameSpace string
}

type Stack struct {
	// Name is the name of the stack
	Name string
	// Services is the number of the services
	ServiceCount int
}

//StackDeploy deploy a new stack
func (client *RolexDockerClient) DeployStack(bundle *Bundle) error {
	//networks := client.getUniqueNetworkNames(bundle.Services)

	//if err := client.updateNetworks(networks, namespace); err != nil {
	//	return err
	//}

	return client.deployServices(bundle.Services, bundle.NameSpace)
}

// StackList list all stack
func (client *RolexDockerClient) ListStack() ([]Stack, error) {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace)
	services, err := client.ListService(types.ServiceListOptions{Filter: filter})
	if err != nil {
		return nil, err
	}

	stackMap := make(map[string]Stack, 0)
	for _, service := range services {
		labels := service.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok {
			log.Errorf("Cannot get label %s for service %s", labelNamespace, service.ID)
			continue
		}

		stack, ok := stackMap[name]
		if !ok {
			stackMap[name] = Stack{
				Name:         name,
				ServiceCount: 1,
			}
		} else {
			stack.ServiceCount++
		}
	}

	var stacks []Stack
	for _, stack := range stackMap {
		stacks = append(stacks, stack)
	}

	return stacks, nil
}

// Inspect stack get stack info
func (client *RolexDockerClient) InspectStack(namespace string) (*Bundle, error) {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace)
	services, err := client.ListService(types.ServiceListOptions{Filter: filter})
	if err != nil {
		return nil, err
	}

	stackServices := make(map[string]bundlefile.Service)
	for _, swarmService := range services {
		labels := swarmService.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok || name != namespace {
			log.Errorf("Cannot get label %s for service %s", labelNamespace, swarmService.ID)
			continue
		}

		stackServices[swarmService.Spec.Name] = client.ConvertStackService(swarmService.Spec)
	}

	return &Bundle{
		NameSpace: namespace,
		Bundlefile: bundlefile.Bundlefile{
			Services: stackServices,
		},
		//TODO stack version is missing
	}, nil
}

// convert swarm service to bundle service
func (client *RolexDockerClient) ConvertStackService(swarmService swarm.ServiceSpec) bundlefile.Service {
	containerSepc := swarmService.TaskTemplate.ContainerSpec
	bundleService := bundlefile.Service{
		Image:      containerSepc.Image,
		Command:    containerSepc.Command,
		Args:       containerSepc.Args,
		Env:        containerSepc.Env,
		WorkingDir: &containerSepc.Dir,
		User:       &containerSepc.User,
		Labels:     containerSepc.Labels,
	}

	var ports []bundlefile.Port
	for _, portSepc := range swarmService.EndpointSpec.Ports {
		ports = append(ports, bundlefile.Port{
			Protocol: string(portSepc.Protocol),
			Port:     portSepc.TargetPort,
		})
	}

	bundleService.Ports = ports
	return bundleService
}

func (client *RolexDockerClient) getUniqueNetworkNames(services map[string]bundlefile.Service) []string {
	networkSet := make(map[string]bool)

	for _, service := range services {
		for _, network := range service.Networks {
			networkSet[network] = true
		}
	}

	networks := []string{}
	for network := range networkSet {
		networks = append(networks, network)
	}

	return networks
}

// update network
func (client *RolexDockerClient) updateNetworks(networks []string, namespace string) error {
	existingNetworks, err := client.filterStackNetwork(namespace)
	if err != nil {
		return err
	}

	existingNetworkMap := make(map[string]goclient.Network)
	for _, network := range existingNetworks {
		existingNetworkMap[network.Name] = network
	}

	labels := client.getStackLabels(namespace, nil)
	createOpts := &goclient.CreateNetworkOptions{
		Options: client.getStackLabelsInterface(labels),
		Driver:  defaultNetworkDriver,
		// docker TODO: remove when engine-api uses omitempty for IPAM
		IPAM: goclient.IPAMOptions{Driver: "default"},
	}

	for _, internalName := range networks {
		name := fmt.Sprintf("%s_%s", namespace, internalName)
		if _, exists := existingNetworkMap[name]; exists {
			continue
		}
		log.Infof("Creating network %s\n", name)

		createOpts.Name = name
		if _, err := client.CreateNetwork(*createOpts); err != nil {
			return err
		}
	}

	return nil
}

func (client *RolexDockerClient) convertNetworks(networks []string, namespace string, name string) []swarm.NetworkAttachmentConfig {
	nets := []swarm.NetworkAttachmentConfig{}
	for _, network := range networks {
		nets = append(nets, swarm.NetworkAttachmentConfig{
			Target:  namespace + "_" + network,
			Aliases: []string{name},
		})
	}

	return nets
}

func (client *RolexDockerClient) deployServices(services map[string]bundlefile.Service, namespace string) error {
	existingServices, err := client.filterStackServices(namespace)
	if err != nil {
		return err
	}

	existingServiceMap := make(map[string]swarm.Service)
	for _, service := range existingServices {
		existingServiceMap[service.Spec.Name] = service
	}

	for internalName, service := range services {
		name := fmt.Sprintf("%s_%s", namespace, internalName)

		var ports []swarm.PortConfig
		for _, portSepc := range service.Ports {
			ports = append(ports, swarm.PortConfig{
				Protocol:   swarm.PortConfigProtocol(portSepc.Protocol),
				TargetPort: portSepc.Port,
			})
		}

		serviceSpec := swarm.ServiceSpec{
			Annotations: swarm.Annotations{
				Name:   name,
				Labels: client.getStackLabels(namespace, service.Labels),
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: swarm.ContainerSpec{
					Image:   service.Image,
					Command: service.Command,
					Args:    service.Args,
					Env:     service.Env,
				},
			},
			EndpointSpec: &swarm.EndpointSpec{
				Ports: ports,
			},
			//Networks: client.convertNetworks(service.Networks, namespace, internalName),
		}

		cspec := &serviceSpec.TaskTemplate.ContainerSpec
		if service.WorkingDir != nil {
			cspec.Dir = *service.WorkingDir
		}

		if service.User != nil {
			cspec.User = *service.User
		}

		if service, exists := existingServiceMap[name]; exists {
			log.Infof("Updating service %s (id %s)", name, service.ID)

			// docker TODO(nishanttotla): Pass headers with X-Registry-Auth
			if err := client.UpdateService(service.ID, service.Version, serviceSpec, nil); err != nil {
				return err
			}
		} else {
			log.Infof("Creating service %s", name)
			// docker TODO(nishanttotla): Pass headers with X-Registry-Auth
			if _, err := client.CreateService(serviceSpec, types.ServiceCreateOptions{}); err != nil {
				return err
			}
		}
	}

	return nil
}

// get stack labels
func (client *RolexDockerClient) getStackLabels(namespace string, labels map[string]string) map[string]string {
	if labels == nil {
		labels = make(map[string]string)
	}

	labels[labelNamespace] = namespace
	return labels
}

// convert labels map[string]string to map[string]interface{}
func (client *RolexDockerClient) getStackLabelsInterface(labels map[string]string) map[string]interface{} {
	labelsInterface := make(map[string]interface{})
	for key, value := range labels {
		labelsInterface[key] = value
	}

	return labelsInterface
}

// split joint stack filter
func (client *RolexDockerClient) getStackFilter(namespace string) filters.Args {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace+"="+namespace)
	return filter
}

// get service by default stack labels
func (client *RolexDockerClient) filterStackServices(namespace string) ([]swarm.Service, error) {
	return client.ListService(types.ServiceListOptions{Filter: client.getStackFilter(namespace)})
}

// get network by default filter
func (client *RolexDockerClient) filterStackNetwork(namespace string) ([]goclient.Network, error) {
	return client.ListNetworks(goclient.NetworkFilterOpts{})
}
