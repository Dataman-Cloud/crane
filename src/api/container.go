package api

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	log "github.com/Sirupsen/logrus"
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
	CONTAINER_RM   = "rm"
)

const (
	CONTAINER_STOP_TIMEOUT = 1 << 20
)

func (api *Api) InspectContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	container, err := api.GetDockerClient().InspectContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	api.HttpResponse(ctx, err, container)
	return
}

func (api *Api) ListContainers(ctx *gin.Context) {
	all, err := strconv.ParseBool(ctx.DefaultQuery("all", "true"))
	if err != nil {
		log.Error("Parse param all of list container got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	size, err := strconv.ParseBool(ctx.DefaultQuery("size", "false"))
	if err != nil {
		log.Error("Parse param size of list container got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	limitValue, err := strconv.ParseInt(ctx.DefaultQuery("limit", "0"), 10, 64)
	if err != nil {
		log.Error("Parse param all of limit container got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}
	limit := int(limitValue)

	filters := make(map[string][]string)
	queryFilters := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(queryFilters), &filters); err != nil {
		log.Error("Unmarshal list container filters got error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	listOpts := goclient.ListContainersOptions{
		All:     all,
		Size:    size,
		Limit:   limit,
		Since:   ctx.DefaultQuery("since", ""),
		Before:  ctx.DefaultQuery("before", ""),
		Filters: filters,
	}

	rolexContext, _ := ctx.Get("rolexContext")
	containers, err := api.GetDockerClient().ListContainers(rolexContext.(context.Context), listOpts)
	api.HttpResponse(ctx, err, containers)
	return
}

func (api *Api) PatchContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	containerRequest := &ContainerRequest{}
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodePatchContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	switch strings.ToLower(containerRequest.Method) {
	case "rename":
		opts := goclient.RenameContainerOptions{
			Name: containerRequest.Name,
			ID:   ctx.Param("container_id"),
		}
		err = api.GetDockerClient().RenameContainer(rolexContext.(context.Context), opts)
	case "stop":
		err = api.GetDockerClient().StopContainer(rolexContext.(context.Context), ctx.Param("container_id"), CONTAINER_STOP_TIMEOUT)
	case "start":
		err = api.GetDockerClient().StartContainer(rolexContext.(context.Context), ctx.Param("container_id"), nil)
	case "restart":
		err = api.GetDockerClient().RestartContainer(rolexContext.(context.Context), ctx.Param("container_id"), CONTAINER_STOP_TIMEOUT)
	case "pause":
		err = api.GetDockerClient().PauseContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	case "unpause":
		err = api.GetDockerClient().UnpauseContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	case "resizetty":
		err = api.GetDockerClient().ResizeContainerTTY(rolexContext.(context.Context), ctx.Param("container_id"), containerRequest.Height, containerRequest.Width)
	default:
		err = rolexerror.NewRolexError(rolexerror.CodePatchContainerMethodUndefined, containerRequest.Method)
	}

	api.HttpResponse(ctx, err, "success")
	return
}

func (api *Api) DeleteContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	containerRequest := &ContainerRequest{}
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodeDeleteContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	if containerRequest.Method == CONTAINER_RM {
		opts := goclient.RemoveContainerOptions{ID: ctx.Param("container_id"), Force: true}
		err = api.GetDockerClient().RemoveContainer(rolexContext.(context.Context), opts)
	} else if containerRequest.Method == CONTAINER_KILL {
		opts := goclient.KillContainerOptions{ID: ctx.Param("container_id")}
		err = api.GetDockerClient().KillContainer(rolexContext.(context.Context), opts)
	} else {
		err = rolexerror.NewRolexError(rolexerror.CodeDeleteContainerMethodUndefined, containerRequest.Method)
	}
	api.HttpResponse(ctx, err, "success")
	return
}

func (api *Api) DiffContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	changes, err := api.GetDockerClient().DiffContainer(rolexContext.(context.Context), ctx.Param("container_id"))
	api.HttpResponse(ctx, err, changes)
	return
}

func (api *Api) LogsContainer(ctx *gin.Context) {
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

func (api *Api) StatsContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	stats := make(chan *model.ContainerStat)

	defer close(stats)

	go api.GetDockerClient().StatsContainer(rolexContext.(context.Context), ctx.Param("container_id"), stats)

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "container-stats",
			Data:  <-stats,
		}.Render(ctx.Writer)
		return true
	})
}
