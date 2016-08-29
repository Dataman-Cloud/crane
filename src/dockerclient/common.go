package dockerclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/docker/engine-api/types/swarm"

	rauth "github.com/Dataman-Cloud/go-component/registryauth"
	docker "github.com/Dataman-Cloud/go-dockerclient"
)

// convert swarm service to bundle service
func (client *RolexDockerClient) ToRolexServiceSpec(swarmService swarm.ServiceSpec) model.RolexServiceSpec {
	networks := client.getServiceNetworkNames(swarmService.Networks)
	rolexServiceSpec := model.RolexServiceSpec{
		Name:         swarmService.Name,
		Labels:       swarmService.Labels,
		TaskTemplate: swarmService.TaskTemplate,
		Mode:         swarmService.Mode,
		Networks:     networks,
		UpdateConfig: swarmService.UpdateConfig,
		EndpointSpec: swarmService.EndpointSpec,
	}

	if rolexServiceSpec.UpdateConfig == nil {
		rolexServiceSpec.UpdateConfig = &swarm.UpdateConfig{}
	}

	if rolexServiceSpec.EndpointSpec == nil {
		rolexServiceSpec.EndpointSpec = &swarm.EndpointSpec{}
	}

	if rolexServiceSpec.Labels != nil {
		if registryauth, ok := rolexServiceSpec.Labels[LabelRegistryAuth]; ok {
			rolexServiceSpec.RegistryAuth = registryauth
		}
	}

	return rolexServiceSpec
}

func EncodedRegistryAuth(registryAuth string) (string, error) {
	authInfo, err := rauth.GetHubApi().Get(registryAuth)
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
