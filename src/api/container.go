package api

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	"github.com/Dataman-Cloud/crane/src/utils/dmgin"
	"github.com/Dataman-Cloud/crane/src/utils/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
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

const (
	CodeListContainerParamError        = "400-11001"
	CodePatchContainerParamError       = "400-11002"
	CodePatchContainerMethodUndefined  = "400-11003"
	CodeDeleteContainerParamError      = "400-11004"
	CodeDeleteContainerMethodUndefined = "400-11005"
)

func (api *Api) InspectContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	cId := ctx.Param("container_id")
	container, err := api.GetDockerClient().InspectContainer(rolexContext.(context.Context), cId)
	if err != nil {
		log.Errorf("InspectContainer of containerId %s got error: %s", cId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, container)
	return
}

func (api *Api) ListContainers(ctx *gin.Context) {
	all, err := strconv.ParseBool(ctx.DefaultQuery("all", "true"))
	if err != nil {
		log.Error("Parse param all of list container got error: ", err)
		rerror := rolexerror.NewError(CodeListContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	size, err := strconv.ParseBool(ctx.DefaultQuery("size", "false"))
	if err != nil {
		log.Error("Parse param size of list container got error: ", err)
		rerror := rolexerror.NewError(CodeListContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	limitValue, err := strconv.ParseInt(ctx.DefaultQuery("limit", "0"), 10, 64)
	if err != nil {
		log.Error("Parse param all of limit container got error: ", err)
		rerror := rolexerror.NewError(CodeListContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}
	limit := int(limitValue)

	filters := make(map[string][]string)
	queryFilters := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(queryFilters), &filters); err != nil {
		log.Error("Unmarshal list container filters got error: ", err)
		rerror := rolexerror.NewError(CodeListContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	listOpts := docker.ListContainersOptions{
		All:     all,
		Size:    size,
		Limit:   limit,
		Since:   ctx.DefaultQuery("since", ""),
		Before:  ctx.DefaultQuery("before", ""),
		Filters: filters,
	}

	rolexContext, _ := ctx.Get("rolexContext")
	containers, err := api.GetDockerClient().ListContainers(rolexContext.(context.Context), listOpts)
	if err != nil {
		log.Error("ListContainers got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, containers)
	return
}

func (api *Api) PatchContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var containerRequest ContainerRequest
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewError(CodePatchContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	method := strings.ToLower(containerRequest.Method)
	cId := ctx.Param("container_id")
	switch method {
	case "rename":
		opts := docker.RenameContainerOptions{
			Name: containerRequest.Name,
			ID:   cId,
		}
		err = api.GetDockerClient().RenameContainer(rolexContext.(context.Context), opts)
	case "stop":
		err = api.GetDockerClient().StopContainer(rolexContext.(context.Context), cId, CONTAINER_STOP_TIMEOUT)
	case "start":
		err = api.GetDockerClient().StartContainer(rolexContext.(context.Context), cId, nil)
	case "restart":
		err = api.GetDockerClient().RestartContainer(rolexContext.(context.Context), cId, CONTAINER_STOP_TIMEOUT)
	case "pause":
		err = api.GetDockerClient().PauseContainer(rolexContext.(context.Context), cId)
	case "unpause":
		err = api.GetDockerClient().UnpauseContainer(rolexContext.(context.Context), cId)
	case "resizetty":
		err = api.GetDockerClient().ResizeContainerTTY(rolexContext.(context.Context), cId, containerRequest.Height, containerRequest.Width)
	default:
		err = rolexerror.NewError(CodePatchContainerMethodUndefined, containerRequest.Method)
	}

	if err != nil {
		log.Errorf("%s container of %s got error: %s", method, cId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) DeleteContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var containerRequest ContainerRequest
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewError(CodeDeleteContainerParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	method := containerRequest.Method
	cId := ctx.Param("container_id")
	if method == CONTAINER_RM {
		opts := docker.RemoveContainerOptions{ID: cId, Force: true}
		err = api.GetDockerClient().RemoveContainer(rolexContext.(context.Context), opts)
	} else if method == CONTAINER_KILL {
		opts := docker.KillContainerOptions{ID: cId}
		err = api.GetDockerClient().KillContainer(rolexContext.(context.Context), opts)
	} else {
		err = rolexerror.NewError(CodeDeleteContainerMethodUndefined, containerRequest.Method)
	}

	if err != nil {
		log.Errorf("%s container of %s got error %s", method, cId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) DiffContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	cId := ctx.Param("container_id")
	changes, err := api.GetDockerClient().DiffContainer(rolexContext.(context.Context), cId)
	if err != nil {
		log.Errorf("Diff container of %s got error: %s", cId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, changes)
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

	chnMsg := make(chan *model.RolexContainerStat)
	defer close(chnMsg)
	chnDone := make(chan bool)
	defer close(chnDone)
	chnContainerStats := make(chan *docker.Stats) // closed by go-dockerclient

	chnErr := make(chan error, 1)
	defer close(chnErr)

	cId := ctx.Param("container_id")
	opts := model.ContainerStatOptions{
		ID:                  cId,
		Stats:               chnContainerStats,
		Stream:              true,
		Done:                chnDone,
		RolexContainerStats: chnMsg,
	}

	go func() {
		chnErr <- api.GetDockerClient().StatsContainer(rolexContext.(context.Context), opts)
	}()

	ssEvent := &sse.Event{Event: "container-stats"}
	w := ctx.Writer
	clientGone := w.CloseNotify()
	var clientClosed bool = false
	for {
		select {
		case <-clientGone:
			clientClosed = true
			log.Infof("Stats stream of container %s closed by client", cId)
			chnDone <- true
		case data := <-chnMsg:
			if !clientClosed {
				ssEvent.Data = data
				ssEvent.Render(w)
				w.Flush()
			}
		case err := <-chnErr:
			log.Errorf("Stats container of %s stop with error: %s", cId, err)
			return
		}
	}
}
