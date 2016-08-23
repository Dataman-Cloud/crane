package dockerclient

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

const (
	labelNamespace = "com.docker.stack.namespace"
)

type Stacks []Stack

func (s Stacks) Len() int {
	return len(s)
}

func (s Stacks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Stacks) Less(i, j int) bool {
	return s[j].Services[0].CreatedAt.Unix() < s[i].Services[0].CreatedAt.Unix()
}

type Stack struct {
	// Name is the name of the stack
	Namespace string `json:"Namespace"`
	// Services is the number of the services
	ServiceCount int `json:"ServiceCount"`

	Services []ServiceStatus
}

// deploy a new stack
func (client *RolexDockerClient) DeployStack(bundle *model.Bundle) error {
	if bundle.Namespace == "" || !isValidName.MatchString(bundle.Namespace) {
		return rolexerror.NewRolexError(rolexerror.CodeInvalidStackName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]")
	}

	networks := client.getUniqueNetworkNames(bundle.Stack.Services)

	newNetworkMap, err := client.updateNetworks(networks, bundle.Namespace)
	if err != nil {
		return err
	}

	return client.deployServices(bundle.Stack.Services, bundle.Namespace, newNetworkMap)
}

// list all stack
func (client *RolexDockerClient) ListStack() ([]Stack, error) {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace)
	services, err := client.ListServiceSpec(types.ServiceListOptions{Filter: filter})
	if err != nil {
		return nil, err
	}

	stackMap := make(map[string]*Stack, 0)
	for _, service := range services {
		labels := service.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok {
			log.Warnf("Cannot get label %s for service %s", labelNamespace, service.ID)
			continue
		}

		stack, ok := stackMap[name]
		if !ok {
			stackMap[name] = &Stack{
				Namespace:    name,
				ServiceCount: 1,
			}
		} else {
			stack.ServiceCount++
		}
	}

	var stacks []Stack
	for _, stack := range stackMap {
		stackServices, err := client.ListStackService(stack.Namespace, types.ServiceListOptions{})
		if err == nil {
			stack.Services = stackServices
		}
		stacks = append(stacks, *stack)
	}
	sort.Sort(Stacks(stacks))

	return stacks, nil
}

// ListStackServices return list of service staus and core config in stack
func (client *RolexDockerClient) ListStackService(namespace string, opts types.ServiceListOptions) ([]ServiceStatus, error) {
	services, err := client.FilterServiceByStack(namespace, opts)
	if err != nil {
		return nil, err
	}

	return client.GetServicesStatus(services)
}

// Inspect stack get stack info
func (client *RolexDockerClient) InspectStack(namespace string) (*model.Bundle, error) {
	services, err := client.FilterServiceByStack(namespace, types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}

	stackServices := make(map[string]model.RolexService)
	for _, swarmService := range services {
		stackServices[swarmService.Spec.Name] = client.ConvertStackService(swarmService.Spec)
	}

	return &model.Bundle{
		Namespace: namespace,
		Stack: model.BundleService{
			//TODO stack version is missing
			Services: stackServices,
		},
	}, nil
}

// remove all service and network in the stack
func (client *RolexDockerClient) RemoveStack(namespace string) error {
	services, err := client.FilterServiceByStack(namespace, types.ServiceListOptions{})
	if err != nil {
		return err
	}

	for _, service := range services {
		log.Info("begin to remove service ", service.Spec.Name)
		if err := client.RemoveService(service.ID); err != nil {
			return err
		}
	}

	networks, err := client.filterStackNetwork(namespace)
	if err != nil {
		return err
	}

	for _, network := range networks {
		log.Info("begin to remove network ", network.Name)
		if err := client.RemoveNetwork(network.ID); err != nil {
			return err
		}
	}

	if len(services) == 0 && len(networks) == 0 {
		return rolexerror.NewRolexError(rolexerror.CodeStackNotFound, fmt.Sprintf("stack %s not found", namespace))
	}

	return nil
}

// filter service by stack name
func (client *RolexDockerClient) FilterServiceByStack(namespace string, opts types.ServiceListOptions) ([]swarm.Service, error) {
	if opts.Filter.Len() == 0 {
		opts.Filter = filters.NewArgs()
	}
	opts.Filter.Add("label", labelNamespace)
	services, err := client.ListServiceSpec(opts)
	if err != nil {
		return nil, err
	}

	var stackServices []swarm.Service
	for _, service := range services {
		labels := service.Spec.Labels
		name, ok := labels[labelNamespace]
		if !ok {
			log.Warnf("Cannot get label %s for service %s", labelNamespace, service.ID)
			continue
		}

		if name != namespace {
			continue
		}

		stackServices = append(stackServices, service)
	}

	return stackServices, nil
}

func (client *RolexDockerClient) GetStackGroup(namespace string) (uint64, error) {
	bundle, err := client.InspectStack(namespace)
	if err != nil {
		return 0, err
	}

	for _, service := range bundle.Stack.Services {
		for k, _ := range service.Labels {
			if strings.HasPrefix(k, "com.rolex.permissions") {
				groupId, err := strconv.ParseUint(strings.Split(k, ".")[3], 10, 64)
				if err == nil {
					return groupId, nil
				}
			}
		}
	}
	return 0, errors.New("can't found stack groupid")
}

// convert swarm service to bundle service
func (client *RolexDockerClient) ConvertStackService(swarmService swarm.ServiceSpec) model.RolexService {
	networks := client.getServiceNetworks(swarmService.Networks)
	return model.RolexService{
		Name:         swarmService.Name,
		Labels:       swarmService.Labels,
		TaskTemplate: swarmService.TaskTemplate,
		Mode:         swarmService.Mode,
		UpdateConfig: swarmService.UpdateConfig,
		Networks:     networks,
		EndpointSpec: swarmService.EndpointSpec,
	}
}

func (client *RolexDockerClient) getUniqueNetworkNames(services map[string]model.RolexService) []string {
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
func (client *RolexDockerClient) updateNetworks(networks []string, namespace string) (map[string]bool, error) {
	existingNetworks, err := client.ListNetworks(docker.NetworkFilterOpts{})
	if err != nil {
		return nil, err
	}

	existingNetworkMap := make(map[string]docker.Network)
	for _, network := range existingNetworks {
		existingNetworkMap[network.Name] = network
	}

	createOpts := &docker.CreateNetworkOptions{
		Labels: client.getStackLabels(namespace, nil),
		Driver: defaultNetworkDriver,
		// docker TODO: remove when engine-api uses omitempty for IPAM
		IPAM: docker.IPAMOptions{Driver: "default"},
	}

	newNetworkMap := make(map[string]bool)
	for _, internalName := range networks {
		if _, exists := existingNetworkMap[internalName]; exists {
			newNetworkMap[internalName] = false
			continue
		}

		name := fmt.Sprintf("%s_%s", namespace, internalName)
		log.Infof("Creating network %s\n", name)
		createOpts.Name = name
		if _, err := client.CreateNetwork(*createOpts); err != nil {
			return newNetworkMap, err
		}
		newNetworkMap[name] = true
	}

	return newNetworkMap, nil
}

func (client *RolexDockerClient) convertNetworks(newNetworkMap map[string]bool, namespace string, name string) []swarm.NetworkAttachmentConfig {
	nets := []swarm.NetworkAttachmentConfig{}
	for network, isNew := range newNetworkMap {
		if isNew {
			nets = append(nets, swarm.NetworkAttachmentConfig{
				Target:  namespace + "_" + network,
				Aliases: []string{name},
			})

		} else {
			nets = append(nets, swarm.NetworkAttachmentConfig{
				Target:  network,
				Aliases: []string{name},
			})
		}
	}

	return nets
}

func (client *RolexDockerClient) getServiceNetworks(nets []swarm.NetworkAttachmentConfig) []string {
	networkList := []string{}
	for _, network := range nets {
		networkList = append(networkList, network.Target)
	}

	return networkList
}

func (client *RolexDockerClient) deployServices(services map[string]model.RolexService, namespace string, newNetworkMap map[string]bool) error {
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
		serviceSpec := swarm.ServiceSpec{
			Annotations: swarm.Annotations{
				Name:   name,
				Labels: client.getStackLabels(namespace, service.Labels),
			},
			Mode:         service.Mode,
			TaskTemplate: service.TaskTemplate,
			EndpointSpec: service.EndpointSpec,
			Networks:     client.convertNetworks(newNetworkMap, namespace, internalName),
			UpdateConfig: service.UpdateConfig,
		}

		//TODO change service WorkingDir and User
		//cspec := &serviceSpec.TaskTemplate.ContainerSpec
		//if service.WorkingDir != nil {
		//	cspec.Dir = *service.WorkingDir
		//}

		//if service.User != nil {
		//	cspec.User = *service.User
		//}

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

// split joint stack filter
func (client *RolexDockerClient) getStackFilter(namespace string) filters.Args {
	filter := filters.NewArgs()
	filter.Add("label", labelNamespace+"="+namespace)
	return filter
}

// get service by default stack labels
func (client *RolexDockerClient) filterStackServices(namespace string) ([]swarm.Service, error) {
	return client.ListServiceSpec(types.ServiceListOptions{Filter: client.getStackFilter(namespace)})
}

// get network by default filter
func (client *RolexDockerClient) filterStackNetwork(namespace string) ([]docker.Network, error) {
	filter := docker.NetworkFilterOpts{"label": map[string]bool{labelNamespace: true}}
	networks, err := client.ListNetworks(filter)
	if err != nil {
		return nil, err
	}

	var stackNetwork []docker.Network
	for _, network := range networks {
		if network.Labels == nil {
			continue
		}

		if name, ok := network.Labels[labelNamespace]; !ok || name != namespace {
			continue
		}

		stackNetwork = append(stackNetwork, network)
	}

	return stackNetwork, nil
}
