package api

import (
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeCreateVolumeParamError = "400-11601"
)

func (api *Api) InspectVolume(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	volume, err := api.GetDockerClient().InspectVolume(craneContext.(context.Context), ctx.Param("volume_id"))
	if err != nil {
		log.Errorf("inspect volume error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, volume)
	return
}

func (api *Api) ListVolume(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	var opts docker.ListVolumesOptions

	volumes, err := api.GetDockerClient().ListVolumes(craneContext.(context.Context), opts)
	if err != nil {
		log.Errorf("get volume list error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, volumes)
	return
}

func (api *Api) CreateVolume(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	var opts docker.CreateVolumeOptions

	if err := ctx.BindJSON(&opts); err != nil {
		log.Errorf("create volume request body parse json error: %v", err)
		craneError := cranerror.NewError(CodeCreateVolumeParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	volume, err := api.GetDockerClient().CreateVolume(craneContext.(context.Context), opts)
	if err != nil {
		log.Errorf("create volume error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, volume)
	return
}

func (api *Api) RemoveVolume(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	if err := api.GetDockerClient().RemoveVolume(craneContext.(context.Context), ctx.Param("volume_id")); err != nil {
		log.Errorf("remove volume error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}
