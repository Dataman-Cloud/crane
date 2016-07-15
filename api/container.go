package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/util"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

type ContainerDeleteRequest struct {
	Method string `json:"method"`
}

const (
	CONTAINER_KILL = "kill"
	CONTAINER_RMF  = "rm"
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

func (api *Api) DeleteContainer(ctx *gin.Context) {
	containerDeleteRequest := &ContainerDeleteRequest{}
	if err := ctx.BindJSON(&containerDeleteRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	if containerDeleteRequest.Method == CONTAINER_RMF {
		opts := goclient.RemoveContainerOptions{ID: ctx.Param("container_id"), Force: true}
		err := api.GetDockerClient().RemoveContainer(opts)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0})
	} else if containerDeleteRequest.Method == CONTAINER_KILL {
		opts := goclient.KillContainerOptions{ID: ctx.Param("container_id")}
		err := api.GetDockerClient().KillContainer(opts)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1})
	}
}

func (api *Api) DiffContainer(ctx *gin.Context) {
	changes, err := api.GetDockerClient().DiffContainer(ctx.Param("container_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": changes})
}
