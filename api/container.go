package api

import (
	"net/http"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectContainer(ctx *gin.Context) {
	container, err := api.GetDockerClient().InspectContainer(ctx.Param("container_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": container})
}

func (api *Api) ListContainers(ctx *gin.Context) {
	containers, err := api.GetDockerClient().ListContainers(goclient.ListContainersOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": containers})
}

func (api *Api) CreateContainer(ctx *gin.Context) {}
func (api *Api) UpdateContainer(ctx *gin.Context) {}

func (api *Api) RemoveContainer(ctx *gin.Context) {
	opts := goclient.RemoveContainerOptions{ID: ctx.Param("container_id"), Force: true}
	err := api.GetDockerClient().RemoveContainer(opts)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func (api *Api) KillContainer(ctx *gin.Context) {
	opts := goclient.KillContainerOptions{ID: ctx.Param("container_id")}
	err := api.GetDockerClient().KillContainer(opts)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func (api *Api) DiffContainer(ctx *gin.Context) {
	changes, err := api.GetDockerClient().DiffContainer(ctx.Param("container_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": changes})
}
