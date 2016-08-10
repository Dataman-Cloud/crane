package api

import (
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (api *Api) InspectVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	volume, err := api.GetDockerClient().InspectVolume(rolexContext.(context.Context), ctx.Param("volume_id"))
	if err != nil {
		log.Errorf("inspect volume error: %v", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, volume)
	return
}

func (api *Api) ListVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts docker.ListVolumesOptions

	volumes, err := api.GetDockerClient().ListVolumes(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("get volume list error: %v", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, volumes)
	return
}

func (api *Api) CreateVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var opts docker.CreateVolumeOptions

	if err := ctx.BindJSON(&opts); err != nil {
		log.Errorf("create volume request body parse json error: %v", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateVolumeParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	volume, err := api.GetDockerClient().CreateVolume(rolexContext.(context.Context), opts)
	if err != nil {
		log.Errorf("create volume error: %v", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, volume)
	return
}

func (api *Api) RemoveVolume(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	if err := api.GetDockerClient().RemoveVolume(rolexContext.(context.Context), ctx.Param("volume_id")); err != nil {
		log.Errorf("remove volume error: %v", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}
