package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/util/rerror"

	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectService(ctx *gin.Context) {}
func (api *Api) UpdateService(ctx *gin.Context)  {}

// ServiceCreate creates a new Service.
func (api *Api) CreateService(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	response, err := api.GetDockerClient().CreateService(service, types.ServiceCreateOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": response})
	return
}

// ServiceList returns the list of services.
func (api *Api) ListService(ctx *gin.Context) {
	services, err := api.GetDockerClient().ListService(types.ServiceListOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": services})
	return
}

func (api *Api) RemoveService(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	if serviceID == "" {
		ctx.JSON(http.StatusBadRequest, "service id is nil")
		return
	}

	if err := api.GetDockerClient().RemoveService(serviceID); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
	return
}

func (api *Api) ScaleService(ctx *gin.Context) {
	serviceID := ctx.Param("service_id")
	var serviceScale dockerclient.ServiceScale
	if err := ctx.BindJSON(&serviceScale); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, "requst error"))
		return
	}

	if err := api.GetDockerClient().ScaleService(serviceID, serviceScale); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, "requst error"))
		return
	}

	api.OK(ctx, http.StatusOK, "success")
	return
}
