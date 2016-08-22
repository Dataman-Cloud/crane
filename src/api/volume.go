package api

import (
	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/go-component/utils/dmgin"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeCreateVolumeParamError = "400-11601"
)

func (api *Api) InspectVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	volume, err := api.GetDockerClient().InspectVolume(rolexContext.(context.Context), ctx.Param("volume_id"))
	if err != nil {
		log.Errorf("inspect volume error: %v", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, volume)
	return
}

func (api *Api) ListVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts docker.ListVolumesOptions

	volumes, err := api.GetDockerClient().ListVolumes(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("get volume list error: %v", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, volumes)
	return
}

func (api *Api) CreateVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts docker.CreateVolumeOptions

	if err := ctx.BindJSON(&opts); err != nil {
		log.Errorf("create volume request body parse json error: %v", err)
		rerror := dmerror.NewError(CodeCreateVolumeParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	volume, err := api.GetDockerClient().CreateVolume(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("create volume error: %v", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, volume)
	return
}

func (api *Api) RemoveVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	if err := api.GetDockerClient().RemoveVolume(rolexContext.(context.Context), ctx.Param("volume_id")); err != nil {
		log.Errorf("remove volume error: %v", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}
