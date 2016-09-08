package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/rolex/src/utils/dmgin"
	"github.com/Dataman-Cloud/rolex/src/utils/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeConnectNetworkParamError  = "400-11201"
	CodeConnectNetworkMethodError = "400-11202"
	CodeCreateNetworkParamError   = "400-11203"
	CodeInspectNetworkParamError  = "400-11204"
	CodeListNetworkParamError     = "400-11205"
)

type ConnectNetworkRequest struct {
	Method         string
	NetworkOptions docker.NetworkConnectionOptions
}

const (
	NETWORK_CONNECT    = "connect"
	NETWORK_DISCONNECT = "disconnect"
)

func (api *Api) ConnectNetwork(ctx *gin.Context) {
	var connectNetworkRequest ConnectNetworkRequest

	if err := ctx.BindJSON(&connectNetworkRequest); err != nil {
		log.Errorf("connect network request body parse json error: %v", err)
		rerror := rolexerror.NewError(CodeConnectNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	networkID := ctx.Param("network_id")
	var err error
	switch connectNetworkRequest.Method {
	case NETWORK_CONNECT:
		err = api.GetDockerClient().ConnectNetwork(networkID, connectNetworkRequest.NetworkOptions)
	case NETWORK_DISCONNECT:
		err = api.GetDockerClient().DisconnectNetwork(networkID, connectNetworkRequest.NetworkOptions)
	default:
		err = rolexerror.NewError(CodeConnectNetworkMethodError, connectNetworkRequest.Method)
	}

	if err != nil {
		log.Errorf("disconnect to network %s got error: %s", networkID, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) CreateNetwork(ctx *gin.Context) {
	var netWorkOption docker.CreateNetworkOptions

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		log.Error("create network request body parse json error: ", err)
		rerror := rolexerror.NewError(CodeCreateNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	network, err := api.GetDockerClient().CreateNetwork(netWorkOption)
	if err != nil {
		log.Error("create network error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, network)
	return
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	network, err := api.GetDockerClient().InspectNetwork(ctx.Param("network_id"))
	if err != nil {
		log.Error("inspect network error: ", err)
		rerror := rolexerror.NewError(CodeInspectNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	dmgin.HttpOkResponse(ctx, network)
	return
}

func (api *Api) ListNetworks(ctx *gin.Context) {
	filter := docker.NetworkFilterOpts{}

	fp := ctx.DefaultQuery("filters", "{\"driver\": {\"overlay\": true}}")
	if err := json.Unmarshal([]byte(fp), &filter); err != nil {
		log.Error("list network request body parse json error: ", err)
		rerror := rolexerror.NewError(CodeListNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	networks, err := api.GetDockerClient().ListNetworks(filter)
	if err != nil {
		log.Error("list network get network list error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, networks)
	return
}

func (api *Api) RemoveNetwork(ctx *gin.Context) {
	if err := api.GetDockerClient().RemoveNetwork(ctx.Param("network_id")); err != nil {
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "remove succsee")
	return
}

func (api *Api) ConnectNodeNetwork(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var connectNetworkRequest ConnectNetworkRequest
	if err := ctx.BindJSON(&connectNetworkRequest); err != nil {
		log.Errorf("connect network request body parse json error: %v", err)
		rerror := rolexerror.NewError(CodeConnectNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	nodeID := ctx.Param("node_id")
	networkID := ctx.Param("network_id")
	var err error
	method := connectNetworkRequest.Method
	switch method {
	case NETWORK_CONNECT:
		err = api.GetDockerClient().ConnectNodeNetwork(rolexContext.(context.Context), networkID, connectNetworkRequest.NetworkOptions)
	case NETWORK_DISCONNECT:
		err = api.GetDockerClient().DisconnectNodeNetwork(rolexContext.(context.Context), networkID, connectNetworkRequest.NetworkOptions)
	default:
		err = rolexerror.NewError(CodeConnectNetworkMethodError, method)
	}

	if err != nil {
		log.Errorf("%s to node: %s network %s got error: %s", method, nodeID, networkID, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) InspectNodeNetwork(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	nodeID := ctx.Param("node_id")
	networkID := ctx.Param("network_id")
	network, err := api.GetDockerClient().InspectNodeNetwork(rolexContext.(context.Context), networkID)
	if err != nil {
		log.Errorf("inspect network of node: %s networkid: %s got error: %s", nodeID, networkID, err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, network)
	return
}

func (api *Api) ListNodeNetworks(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	nodeID := ctx.Param("node_id")
	var filters docker.NetworkFilterOpts

	fp := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(fp), &filters); err != nil {
		log.Error("list network request body parse json error: ", err)
		rerror := rolexerror.NewError(CodeListNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	networks, err := api.GetDockerClient().ListNodeNetworks(rolexContext.(context.Context), filters)
	if err != nil {
		log.Errorf("list network get network of %s got error: %s", nodeID, err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, networks)
	return
}

func (api *Api) CreateNodeNetwork(ctx *gin.Context) {
	var netWorkOption docker.CreateNetworkOptions
	rolexContext, _ := ctx.Get("rolexContext")

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		log.Error("create node network request body parse json error: ", err)
		rerror := rolexerror.NewError(CodeCreateNetworkParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	network, err := api.GetDockerClient().CreateNodeNetwork(rolexContext.(context.Context), netWorkOption)
	if err != nil {
		log.Error("create network error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, network)
	return
}
