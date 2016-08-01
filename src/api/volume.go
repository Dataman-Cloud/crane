package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/util"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (api *Api) InspectVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	volume, err := api.GetDockerClient().InspectVolume(rolexContext.(context.Context), ctx.Param("volume_id"))
	if err != nil {
		log.Errorf("inspect volume error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": volume})
}

func (api *Api) ListVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts goclient.ListVolumesOptions

	/*if err := ctx.BindJSON(&opts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}*/

	volumes, err := api.GetDockerClient().ListVolumes(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("get volume list error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": volumes})
}

func (api *Api) CreateVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts goclient.CreateVolumeOptions

	if err := ctx.BindJSON(&opts); err != nil {
		log.Errorf("create volume request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	volume, err := api.GetDockerClient().CreateVolume(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("create volume error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": util.OPERATION_SUCCESS, "data": volume})
}

func (api *Api) RemoveVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	if err := api.GetDockerClient().RemoveVolume(rolexContext.(context.Context), ctx.Param("volume_id")); err != nil {
		log.Errorf("remove volume error: %v", err)
		ctx.JSON(http.StatusForbidden, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": "remove success"})
}
