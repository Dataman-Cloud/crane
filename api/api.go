package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/util"

	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/gin-gonic/gin"
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

func (api *Api) HttpResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": data})
		return
	}

	switch err.(type) {
	case *util.StatusForbiddenError:
		ctx.JSON(http.StatusForbidden, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	default:
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}
}
