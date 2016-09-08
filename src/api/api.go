package api

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"
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
