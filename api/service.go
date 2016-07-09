package api

import (
	"net/http"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectService(ctx *gin.Context) {}
func (api *Api) ListService(ctx *gin.Context)    {}
func (api *Api) UpdateService(ctx *gin.Context)  {}
func (api *Api) RemoveService(ctx *gin.Context)  {}

func (api *Api) ServiceCreate(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	response, err := api.GetDockerClient().ServiceCreate(service, types.ServiceCreateOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": response})
	return
}
