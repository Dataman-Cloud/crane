package api

import (
	"io"
	"net/http"
	"strings"

	"github.com/Dataman-Cloud/rolex/util"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"github.com/manucorporat/sse"
	"golang.org/x/net/context"
)

type ContainerRequest struct {
	Method string `json:"Method"`
	Name   string `json:"Name"`
	Height int    `json:"Height"`
	Width  int    `json:"Width"`
}

const (
	CONTAINER_KILL = "kill"
	CONTAINER_RMF  = "rm"
)

const (
	CONTAINER_STOP_TIMEOUT = 1 << 20
)

func (api *Api) InspectContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	container, err := api.GetDockerClient().InspectContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": container})
}

func (api *Api) ListContainers(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	containers, err := api.GetDockerClient().ListContainers(rolexContext.(context.Context), goclient.ListContainersOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": containers})
}

func (api *Api) PatchContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	containerRequest := &ContainerRequest{}
	if err := ctx.BindJSON(&containerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	switch strings.ToLower(containerRequest.Method) {
	case "rename":
		opts := goclient.RenameContainerOptions{
			Name: containerRequest.Name,
			ID:   ctx.Param("container_id"),
		}
		err := api.GetDockerClient().RenameContainer(rolexContext.(context.Context), opts)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "stop":
		err := api.GetDockerClient().StopContainer(rolexContext.(context.Context), ctx.Param("container_id"), CONTAINER_STOP_TIMEOUT)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "start":
		err := api.GetDockerClient().StartContainer(rolexContext.(context.Context), ctx.Param("container_id"), nil)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "restart":
		err := api.GetDockerClient().RestartContainer(rolexContext.(context.Context), ctx.Param("container_id"), CONTAINER_STOP_TIMEOUT)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "pause":
		err := api.GetDockerClient().PauseContainer(rolexContext.(context.Context), ctx.Param("container_id"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "unpause":
		err := api.GetDockerClient().UnpauseContainer(rolexContext.(context.Context), ctx.Param("container_id"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	case "resizetty":
		err := api.GetDockerClient().ResizeContainerTTY(rolexContext.(context.Context), ctx.Param("container_id"), containerRequest.Height, containerRequest.Width)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": ""})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (api *Api) DeleteContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	containerRequest := &ContainerRequest{}
	if err := ctx.BindJSON(&containerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	if containerRequest.Method == CONTAINER_RMF {
		opts := goclient.RemoveContainerOptions{ID: ctx.Param("container_id"), Force: true}
		err := api.GetDockerClient().RemoveContainer(rolexContext.(context.Context), opts)
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": 0})
	} else if containerRequest.Method == CONTAINER_KILL {
		opts := goclient.KillContainerOptions{ID: ctx.Param("container_id")}
		err := api.GetDockerClient().KillContainer(rolexContext.(context.Context), opts)
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
	rolexContext, _ := ctx.Get("rolexContext")
	changes, err := api.GetDockerClient().DiffContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": changes})
}

func (api *Api) Logs(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	message := make(chan string)

	defer close(message)

	go api.GetDockerClient().LogsContainer(rolexContext.(context.Context), ctx.Param("container_id"), message)

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "container-logs",
			Data:  <-message,
		}.Render(ctx.Writer)
		return true
	})
}

func (api *Api) Stats(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	stats := make(chan *goclient.Stats, 10)
	done := make(chan bool)

	defer func() {
		done <- true
	}()

	go api.GetDockerClient().StatsContainer(rolexContext.(context.Context), ctx.Param("container_id"), stats, done)

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "container-stats",
			Data:  <-stats,
		}.Render(ctx.Writer)
		return true
	})
}
