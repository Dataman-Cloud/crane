package api

import (
	"encoding/json"
	"strconv"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (api *Api) ListImages(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	all, err := strconv.ParseBool(ctx.DefaultQuery("all", "false"))
	if err != nil {
		log.Error("Parse param all of list images got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListImageParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	digests, err := strconv.ParseBool(ctx.DefaultQuery("digests", "true"))
	if err != nil {
		log.Error("Parse param digests of list images got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListImageParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	filters := make(map[string][]string)
	queryFilters := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(queryFilters), &filters); err != nil {
		log.Error("Unmarshal list images filters got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListImageParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	opts := docker.ListImagesOptions{
		All:     all,
		Digests: digests,
		Filter:  ctx.Query("filter"),
		Filters: filters,
	}

	images, err := api.GetDockerClient().ListImages(rolexContext.(context.Context), opts)
	if err != nil {
		log.Error("List images got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, images)
	return
}

func (api *Api) InspectImage(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	image, err := api.GetDockerClient().InspectImage(rolexContext.(context.Context), ctx.Param("image_id"))
	if err != nil {
		log.Error("InspectNetwork got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, image)
	return
}

func (api *Api) ImageHistory(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	historys, err := api.GetDockerClient().ImageHistory(rolexContext.(context.Context), ctx.Param("image_id"))
	if err != nil {
		log.Error("Get ImageHistory got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, historys)
	return
}

// RemoveImage remove image in assign host by image id/name
func (api *Api) RemoveImage(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	imageID := ctx.Param("image_id")
	if err := api.GetDockerClient().RemoveImage(rolexContext.(context.Context), imageID); err != nil {
		log.Error("RemoveImage got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}
