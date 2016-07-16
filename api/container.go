package api

import (
	"io"
	"net/http"

	"github.com/Dataman-Cloud/rolex/util"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"github.com/manucorporat/sse"
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

func (api *Api) Logs(ctx *gin.Context) {
	message := make(chan string)

	defer close(message)

	go api.GetDockerClient().LogsContainer(ctx.Param("node_id"), ctx.Param("container_id"), message)

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "container-logs",
			Data:  <-message,
		}.Render(ctx.Writer)
		return true
	})
}

func (api *Api) Stats(ctx *gin.Context) {
	stats := make(chan *goclient.Stats, 10)
	done := make(chan bool)

	defer func() {
		done <- true
	}()

	go api.GetDockerClient().StatsContainer(ctx.Param("node_id"), ctx.Param("container_id"), stats, done)

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "container-stats",
			Data:  <-stats,
		}.Render(ctx.Writer)
		return true
	})
}
