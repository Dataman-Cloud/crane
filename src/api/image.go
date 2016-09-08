package api

import (
	"encoding/json"
	"strconv"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/dmgin"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	//Image error code
	CodeListImageParamError = "400-11101"
)

func (api *Api) ListImages(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	all, err := strconv.ParseBool(ctx.DefaultQuery("all", "false"))
	if err != nil {
		log.Error("Parse param all of list images got error: ", err)
		rerror := cranerror.NewError(CodeListImageParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	digests, err := strconv.ParseBool(ctx.DefaultQuery("digests", "true"))
	if err != nil {
		log.Error("Parse param digests of list images got error: ", err)
		rerror := cranerror.NewError(CodeListImageParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	filters := make(map[string][]string)
	queryFilters := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(queryFilters), &filters); err != nil {
		log.Error("Unmarshal list images filters got error: ", err)
		rerror := cranerror.NewError(CodeListImageParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	opts := docker.ListImagesOptions{
		All:     all,
		Digests: digests,
		Filter:  ctx.Query("filter"),
		Filters: filters,
	}

	images, err := api.GetDockerClient().ListImages(craneContext.(context.Context), opts)
	if err != nil {
		log.Error("List images got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, images)
	return
}

func (api *Api) InspectImage(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	image, err := api.GetDockerClient().InspectImage(craneContext.(context.Context), ctx.Param("image_id"))
	if err != nil {
		log.Error("InspectNetwork got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, image)
	return
}

func (api *Api) ImageHistory(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	historys, err := api.GetDockerClient().ImageHistory(craneContext.(context.Context), ctx.Param("image_id"))
	if err != nil {
		log.Error("Get ImageHistory got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, historys)
	return
}

// RemoveImage remove image in assign host by image id/name
func (api *Api) RemoveImage(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	imageID := ctx.Param("image_id")
	if err := api.GetDockerClient().RemoveImage(craneContext.(context.Context), imageID); err != nil {
		log.Error("RemoveImage got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}
