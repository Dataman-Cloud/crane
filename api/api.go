package api

import (
	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/Dataman-Cloud/rolex/util/dockerclient"
)

type Api struct {
	Client *dockerclient.RolexDockerClient
	Config *config.Config
}

func (api *Api) GetDockerClient() *dockerclient.RolexDockerClient {
	return api.Client
}

func (api *Api) GetConfig() *config.Config {
	return api.Config
}
