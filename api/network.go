package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dataman-Cloud/rolex/util"

	log "github.com/Sirupsen/logrus"
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
		log.Errorf("connect network request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	networkID := ctx.Param("network_id")
	switch connectNetworkRequest.Method {
	case NETWORK_CONNECT:
		if err := api.GetDockerClient().ConnectNetwork(networkID, connectNetworkRequest.NetworkOptions); err != nil {
			log.Errorf("connect to network %s got error: %s", networkID, err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
			return
		}
	case NETWORK_DISCONNECT:
		if err := api.GetDockerClient().DisconnectNetwork(networkID, connectNetworkRequest.NetworkOptions); err != nil {
			log.Errorf("disconnect to network %s got error: %s", networkID, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
			return
		}
	default:
		log.Error("connect network invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": "Invalid request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": connectNetworkRequest.Method + " success"})
	return
}

func (api *Api) CreateNetwork(ctx *gin.Context) {
	var netWorkOption goclient.CreateNetworkOptions

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		log.Errorf("create network request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	network, err := api.GetDockerClient().CreateNetwork(netWorkOption)
	if err != nil {
		log.Errorf("create network error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": util.OPERATION_SUCCESS, "data": network})
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	network, err := api.GetDockerClient().InspectNetwork(ctx.Param("network_id"))
	if err != nil {
		log.Errorf("inspect network error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": network})
}

func (api *Api) ListNetworks(ctx *gin.Context) {
	filter := goclient.NetworkFilterOpts{}

	fp := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(fp), &filter); err != nil {
		log.Errorf("list network request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	networks, err := api.GetDockerClient().ListNetworks(filter)
	if err != nil {
		log.Errorf("list network get network list error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": networks})
}

func (api *Api) RemoveNetwork(ctx *gin.Context) {
	if err := api.GetDockerClient().RemoveNetwork(ctx.Param("network_id")); err != nil {
		log.Errorf("remove network error: %v", err)
		ctx.JSON(http.StatusForbidden, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": "remove success"})
}

func (api *Api) ConnectNodeNetwork(ctx *gin.Context) {
	var connectNetworkRequest ConnectNetworkRequest
	if err := ctx.BindJSON(&connectNetworkRequest); err != nil {
		log.Errorf("connect network request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	nodeID := ctx.Param("node_id")
	networkID := ctx.Param("network_id")
	switch connectNetworkRequest.Method {
	case NETWORK_CONNECT:
		if err := api.GetDockerClient().ConnectNodeNetwork(nodeID, networkID, connectNetworkRequest.NetworkOptions); err != nil {
			log.Errorf("connect to node: %s network %s got error: %s", nodeID, networkID, err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
			return
		}
	case NETWORK_DISCONNECT:
		if err := api.GetDockerClient().DisconnectNodeNetwork(nodeID, networkID, connectNetworkRequest.NetworkOptions); err != nil {
			log.Errorf("disconnect to node: %s network %s got error: %s", nodeID, networkID, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
			return
		}
	default:
		log.Error("connect network invalid request")
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": "Invalid request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": connectNetworkRequest.Method + " success"})
	return
}

func (api *Api) InspectNodeNetwork(ctx *gin.Context) {
	nodeID := ctx.Param("node_id")
	networkID := ctx.Param("network_id")
	network, err := api.GetDockerClient().InspectNodeNetwork(nodeID, networkID)
	if err != nil {
		log.Errorf("inspect network of node: %s networkid: %s got error: %s", nodeID, networkID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": network})
	return
}

func (api *Api) ListNodeNetworks(ctx *gin.Context) {
	nodeID := ctx.Param("node_id")
	var filters goclient.NetworkFilterOpts

	fp := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(fp), &filters); err != nil {
		log.Error("list network request body parse json error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	networks, err := api.GetDockerClient().ListNodeNetworks(nodeID, filters)
	if err != nil {
		log.Errorf("list network get network of %s got error: %s", nodeID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": networks})
	return
}
