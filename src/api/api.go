package api

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"
)

type Api struct {
	Client *dockerclient.CraneDockerClient
	Config *config.Config
}

func (api *Api) GetDockerClient() *dockerclient.CraneDockerClient {
	return api.Client
}

func (api *Api) GetConfig() *config.Config {
	return api.Config
}
