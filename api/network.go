package api

import (
	"net/http"

	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func (api *Api) ConnectNetwork(ctx *gin.Context) {
	id := ctx.Param("id")
	var netWorkOption goclient.NetworkConnectionOptions
	if err := ctx.BindJSON(&netWorkOption); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}
	if err := api.GetDockerClient().ConnectNetwork(id, netWorkOption); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "connect success"})
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

func (api *Api) DisconnectNetwork(ctx *gin.Context) {
	id := ctx.Param("id")
	var netWorkOption goclient.NetworkConnectionOptions
	if err := ctx.BindJSON(&netWorkOption); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}
	if err := api.GetDockerClient().DisconnectNetwork(id, netWorkOption); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "disconnect success"})
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	id := ctx.Param("id")
	network, err := api.GetDockerClient().InspectNetwork(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": network})
}
func (api *Api) ListNetworks(ctx *gin.Context) {
	var filter goclient.NetworkFilterOpts
	if err := ctx.BindJSON(&filter); err != nil {
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
	id := ctx.Param("id")
	if err := api.GetDockerClient().RemoveNetwork(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"code": 0, "data": "remove success"})
}
