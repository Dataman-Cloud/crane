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
	cId := ctx.Param("container_id")
	container, err := api.GetDockerClient().InspectContainer(rolexContext.(context.Context), cId)
	if err != nil {
		log.Errorf("InspectContainer of containerId %s got error: %s", cId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, container)
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
	if err != nil {
		log.Error("ListContainers got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, containers)
	return
}

func (api *Api) PatchContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var containerRequest ContainerRequest
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodePatchContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	method := strings.ToLower(containerRequest.Method)
	cId := ctx.Param("container_id")
	switch method {
	case "rename":
		opts := goclient.RenameContainerOptions{
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
		err = rolexerror.NewRolexError(rolexerror.CodePatchContainerMethodUndefined, containerRequest.Method)
	}

	if err != nil {
		log.Errorf("%s container of %s got error: %s", method, cId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) DeleteContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var containerRequest ContainerRequest
	if err := ctx.BindJSON(&containerRequest); err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodeDeleteContainerParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	var err error
	method := containerRequest.Method
	cId := ctx.Param("container_id")
	if method == CONTAINER_RM {
		opts := goclient.RemoveContainerOptions{ID: cId, Force: true}
		err = api.GetDockerClient().RemoveContainer(rolexContext.(context.Context), opts)
	} else if method == CONTAINER_KILL {
		opts := goclient.KillContainerOptions{ID: cId}
		err = api.GetDockerClient().KillContainer(rolexContext.(context.Context), opts)
	} else {
		err = rolexerror.NewRolexError(rolexerror.CodeDeleteContainerMethodUndefined, containerRequest.Method)
	}

	if err != nil {
		log.Errorf("%s container of %s got error %s", method, cId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) DiffContainer(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	cId := ctx.Param("container_id")
	changes, err := api.GetDockerClient().DiffContainer(rolexContext.(context.Context), cId)
	if err != nil {
		log.Errorf("Diff container of %s got error: %s", cId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, changes)
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

	chnMsg := make(chan *model.ContainerStat) //close bu rolex docker client
	chnDone := make(chan bool)
	defer close(chnDone)
	chnSt := make(chan *goclient.Stats) // closed by go-dockerclient

	chnErr := make(chan error, 1)
	defer close(chnErr)

	cId := ctx.Param("container_id")
	opts := model.ContainerStatOptions{
		ID:                  cId,
		Stats:               chnSt,
		Stream:              true,
		Done:                chnDone,
		RolexContainerStats: chnMsg,
	}

	go func() {
		chnErr <- api.GetDockerClient().StatsContainer(rolexContext.(context.Context), opts)
	}()

	ssEvent := sse.Event{Event: "container-stats"}
	w := ctx.Writer
	clientGone := w.CloseNotify()
	for {
		select {
		case <-clientGone:
			opts.ClientClosed = true
			log.Infof("Status stream of container %s closed by client", cId)
			chnDone <- true
		case data := <-chnMsg:
			if !opts.ClientClosed {
				ssEvent.Data = data
				ssEvent.Render(w)
			}
		case err := <-chnErr:
			log.Errorf("Status container of %s stop with error: %s", cId, err)
			return
		}
	}
}
