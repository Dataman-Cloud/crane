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

//StackDeploy deploy a new stack
func (client *RolexDockerClient) StackDeploy(bundle *bundlefile.Bundlefile, namespace string) error {
	//networks := client.getUniqueNetworkNames(bundle.Services)

	//if err := client.updateNetworks(networks, namespace); err != nil {
	//	return err
	//}

	return client.deployServices(bundle.Services, namespace)
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
			if err := client.ServiceUpdate(service.ID, service.Version, serviceSpec, nil); err != nil {
				return err
			}
		} else {
			log.Infof("Creating service %s", name)
			// docker TODO(nishanttotla): Pass headers with X-Registry-Auth
			if _, err := client.ServiceCreate(serviceSpec, types.ServiceCreateOptions{}); err != nil {
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
	return client.ServiceList(types.ServiceListOptions{Filter: client.getStackFilter(namespace)})
}

// get network by default filter
func (client *RolexDockerClient) filterStackNetwork(namespace string) ([]goclient.Network, error) {
	return client.ListNetworks(goclient.NetworkFilterOpts{})
}
