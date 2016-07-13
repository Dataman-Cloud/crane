package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dataman-Cloud/rolex/model"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

const (
	NETWORK_CONNECT    = "connect"
	NETWORK_DISCONNECT = "disconnect"
)

func (api *Api) ConnectNetwork(ctx *gin.Context) {
	var connectNetwork model.ConnectNetwork

	if err := ctx.BindJSON(&connectNetwork); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	if connectNetwork.Method == NETWORK_CONNECT {
		if err := api.GetDockerClient().ConnectNetwork(ctx.Param("id"), connectNetwork.NetworkOptions); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
			return
		}
	} else if connectNetwork.Method == NETWORK_DISCONNECT {
		if err := api.GetDockerClient().DisconnectNetwork(ctx.Param("id"), connectNetwork.NetworkOptions); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "Invalid request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": connectNetwork.Method + " success"})
}

func (api *Api) CreateNetwork(ctx *gin.Context) {
	var netWorkOption goclient.CreateNetworkOptions

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	network, err := api.GetDockerClient().CreateNetwork(netWorkOption)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": network})
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	network, err := api.GetDockerClient().InspectNetwork(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": network})
}

func (api *Api) ListNetworks(ctx *gin.Context) {
	filter := goclient.NetworkFilterOpts{}

	fp := ctx.Query("filters")
	if fp == "" {
		fp = "{}"
	}
	if err := json.Unmarshal([]byte(fp), &filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	networks, err := api.GetDockerClient().ListNetworks(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": networks})
}

func (api *Api) RemoveNetwork(ctx *gin.Context) {
	if err := api.GetDockerClient().RemoveNetwork(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "remove success"})
}
