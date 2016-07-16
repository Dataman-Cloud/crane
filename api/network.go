package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dataman-Cloud/rolex/util/rerror"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

type ConnectNetworkRequest struct {
	Method         string
	NetworkOptions goclient.NetworkConnectionOptions
}

const (
	NETWORK_CONNECT    = "connect"
	NETWORK_DISCONNECT = "disconnect"
)

func (api *Api) ConnectNetwork(ctx *gin.Context) {
	var connectNetworkRequest ConnectNetworkRequest

	if err := ctx.BindJSON(&connectNetworkRequest); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	if connectNetworkRequest.Method == NETWORK_CONNECT {
		if err := api.GetDockerClient().ConnectNetwork(ctx.Param("id"), connectNetworkRequest.NetworkOptions); err != nil {
			api.ERROR(ctx, err)
			return
		}
	} else if connectNetworkRequest.Method == NETWORK_DISCONNECT {
		if err := api.GetDockerClient().DisconnectNetwork(ctx.Param("id"), connectNetworkRequest.NetworkOptions); err != nil {
			api.ERROR(ctx, err)
			return
		}
	} else {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, "requst error"))
		return
	}

	api.OK(ctx, http.StatusOK, connectNetworkRequest.Method+"success")
}

func (api *Api) CreateNetwork(ctx *gin.Context) {
	var netWorkOption goclient.CreateNetworkOptions

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	network, err := api.GetDockerClient().CreateNetwork(netWorkOption)
	if err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	api.OK(ctx, http.StatusOK, network)
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	network, err := api.GetDockerClient().InspectNetwork(ctx.Param("id"))
	if err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	api.OK(ctx, http.StatusOK, network)
}

func (api *Api) ListNetworks(ctx *gin.Context) {
	filter := goclient.NetworkFilterOpts{}

	fp := ctx.Query("filters")
	if fp == "" {
		fp = "{}"
	}
	if err := json.Unmarshal([]byte(fp), &filter); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	networks, err := api.GetDockerClient().ListNetworks(filter)
	if err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	api.OK(ctx, http.StatusOK, networks)
}

func (api *Api) RemoveNetwork(ctx *gin.Context) {
	if err := api.GetDockerClient().RemoveNetwork(ctx.Param("id")); err != nil {
		api.ERROR(ctx, rerror.NewRolexError(rerror.PARAMETER_ERROR, err.Error()))
		return
	}

	api.OK(ctx, http.StatusOK, "")
}
