package dockerclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	rauth "github.com/Dataman-Cloud/crane/src/plugins/registryauth"
	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/docker/engine-api/types/swarm"

	docker "github.com/Dataman-Cloud/go-dockerclient"
)

// convert swarm service to bundle service
func (client *CraneDockerClient) ToCraneServiceSpec(swarmService swarm.ServiceSpec) model.CraneServiceSpec {
	networks := client.GetServiceNetworkNames(swarmService.Networks)
	craneServiceSpec := model.CraneServiceSpec{
		Name:         swarmService.Name,
		Labels:       swarmService.Labels,
		TaskTemplate: swarmService.TaskTemplate,
		Mode:         swarmService.Mode,
		Networks:     networks,
		UpdateConfig: swarmService.UpdateConfig,
		EndpointSpec: swarmService.EndpointSpec,
	}

	if craneServiceSpec.UpdateConfig == nil {
		craneServiceSpec.UpdateConfig = &swarm.UpdateConfig{}
	}

	if craneServiceSpec.EndpointSpec == nil {
		craneServiceSpec.EndpointSpec = &swarm.EndpointSpec{}
	}

	if craneServiceSpec.Labels != nil {
		if registryauth, ok := craneServiceSpec.Labels[LabelRegistryAuth]; ok {
			craneServiceSpec.RegistryAuth = registryauth
		}
	}

	return craneServiceSpec
}

func EncodedRegistryAuth(registryAuth string) (string, error) {
	authInfo, err := rauth.Get(registryAuth)
	if err != nil {
		return "", nil
	}

	authConfig := docker.AuthConfiguration{
		Username: authInfo.Username,
		Password: authInfo.Password,
		Email:    "",
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(authConfig); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf.Bytes()), nil
}

func parseEndpoint(endpoint string) (*url.URL, error) {
	conf := config.GetConfig()
	if !strings.Contains(endpoint, "://") {
		endpoint = conf.DockerEntryScheme + "://" + endpoint
	}

	u, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	if !strings.Contains(u.Host, ":") {
		u.Host = u.Host + ":" + conf.DockerEntryPort
	}

	return u, nil
}

func getAdvertiseAddrByEndpoint(endpoint string) (string, error) {
	u, err := parseEndpoint(endpoint)
	if err != nil {
		return "", err
	}

	return strings.Split(u.Host, ":")[0], nil
}

func GetServicesNamespace(spec swarm.ServiceSpec) string {
	if spec.Annotations.Labels == nil {
		return ""
	}

	return spec.Annotations.Labels[labelNamespace]
}
