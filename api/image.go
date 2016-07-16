package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dataman-Cloud/rolex/util"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func (api *Api) ListImages(ctx *gin.Context) {
	all, err := strconv.ParseBool(ctx.DefaultQuery("all", "false"))
	if err != nil {
		log.Error("Parse param all of list images got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	digests, err := strconv.ParseBool(ctx.DefaultQuery("digests", "true"))
	if err != nil {
		log.Error("Parse param digests of list images got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	filters := make(map[string][]string)
	queryFilters := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(queryFilters), &filters); err != nil {
		log.Error("Unmarshal list images filters got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	opts := goclient.ListImagesOptions{
		All:     all,
		Digests: digests,
		Filter:  ctx.Query("filter"),
		Filters: filters,
	}

	images, err := api.GetDockerClient().ListImages(ctx.Param("node_id"), opts)
	if err != nil {
		log.Error("get images list error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": images})
}

func (api *Api) InspectImage(ctx *gin.Context) {
	image, err := api.GetDockerClient().InspectImage(ctx.Param("node_id"), ctx.Param("image_id"))
	if err != nil {
		log.Error("inspect image error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": image})
}

func (api *Api) ImageHistory(ctx *gin.Context) {
	historys, err := api.GetDockerClient().ImageHistory(ctx.Param("node_id"), ctx.Param("image_id"))
	if err != nil {
		log.Error("get image history error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": historys})
}

// RemoveImage remove image in assign host by image id/name
func (api *Api) RemoveImage(ctx *gin.Context) {
	imageID := ctx.Param("image_id")
	nodeID := ctx.Param("node_id")
	if err := api.GetDockerClient().RemoveImage(nodeID, imageID); err != nil {
		log.Error("Remove images in node %s image %s got error", nodeID, imageID, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": "remove image " + imageID + " success"})
	return
}
