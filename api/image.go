package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dataman-Cloud/rolex/util"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func (api *Api) ListImages(ctx *gin.Context) {
	opts := goclient.ListImagesOptions{}

	if all, err := strconv.ParseBool(ctx.Query("all")); err != nil {
		opts.All = false
	} else {
		opts.All = all
	}

	if digests, err := strconv.ParseBool(ctx.Query("digests")); err != nil {
		opts.Digests = true
	} else {
		opts.Digests = digests
	}

	opts.Filter = ctx.Query("filter")

	filters := make(map[string][]string)
	if err := json.Unmarshal([]byte(ctx.Query("filters")), &filters); err == nil {
		opts.Filters = filters
	}

	images, err := api.GetDockerClient().ListImages(ctx.Param("node_id"), opts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": images})
}

func (api *Api) InspectImage(ctx *gin.Context) {
	image, err := api.GetDockerClient().InspectImage(ctx.Param("node_id"), ctx.Param("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": image})
}

func (api *Api) ImageHistory(ctx *gin.Context) {
	historys, err := api.GetDockerClient().ImageHistory(ctx.Param("node_id"), ctx.Param("name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": historys})
}
