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

// ServiceCreate creates a new Service.
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

// ServiceList returns the list of services.
func (api *Api) ServiceList(ctx *gin.Context) {
	services, err := api.GetDockerClient().ServiceList(types.ServiceListOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": services})
	return
}

func (api *Api) ServiceRemove(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	if serviceID == "" {
		ctx.JSON(http.StatusBadRequest, "service id is nil")
		return
	}

	if err := api.GetDockerClient().ServiceRemove(serviceID); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
	return
}
