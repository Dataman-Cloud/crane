package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
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
		rerror := rolexerror.NewRolexError(rolexerror.CodeConnectNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
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
		err = rolexerror.NewRolexError(rolexerror.CodeConnectNetworkMethodError, connectNetworkRequest.Method)
	}

	if err != nil {
		log.Errorf("disconnect to network %s got error: %s", networkID, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) CreateNetwork(ctx *gin.Context) {
	var netWorkOption goclient.CreateNetworkOptions

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		log.Error("create network request body parse json error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	network, err := api.GetDockerClient().CreateNetwork(netWorkOption)
	if err != nil {
		log.Error("create network error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, network)
	return
}

func (api *Api) InspectNetwork(ctx *gin.Context) {
	network, err := api.GetDockerClient().InspectNetwork(ctx.Param("network_id"))
	if err != nil {
		log.Error("inspect network error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeInspectNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	api.HttpOkResponse(ctx, network)
	return
}

func (api *Api) ListNetworks(ctx *gin.Context) {
	filter := goclient.NetworkFilterOpts{}

	fp := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(fp), &filter); err != nil {
		log.Error("list network request body parse json error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	networks, err := api.GetDockerClient().ListNetworks(filter)
	if err != nil {
		log.Error("list network get network list error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, networks)
	return
}

func (api *Api) RemoveNetwork(ctx *gin.Context) {
	if err := api.GetDockerClient().RemoveNetwork(ctx.Param("network_id")); err != nil {
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "remove succsee")
	return
}

func (api *Api) ConnectNodeNetwork(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	var connectNetworkRequest ConnectNetworkRequest
	if err := ctx.BindJSON(&connectNetworkRequest); err != nil {
		log.Errorf("connect network request body parse json error: %v", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeConnectNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
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
		err = rolexerror.NewRolexError(rolexerror.CodeConnectNetworkMethodError, method)
	}

	if err != nil {
		log.Errorf("%s to node: %s network %s got error: %s", method, nodeID, networkID, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) InspectNodeNetwork(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	nodeID := ctx.Param("node_id")
	networkID := ctx.Param("network_id")
	network, err := api.GetDockerClient().InspectNodeNetwork(rolexContext.(context.Context), networkID)
	if err != nil {
		log.Errorf("inspect network of node: %s networkid: %s got error: %s", nodeID, networkID, err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, network)
	return
}

func (api *Api) ListNodeNetworks(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	nodeID := ctx.Param("node_id")
	var filters goclient.NetworkFilterOpts

	fp := ctx.DefaultQuery("filters", "{}")
	if err := json.Unmarshal([]byte(fp), &filters); err != nil {
		log.Error("list network request body parse json error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeListNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	networks, err := api.GetDockerClient().ListNodeNetworks(rolexContext.(context.Context), filters)
	if err != nil {
		log.Errorf("list network get network of %s got error: %s", nodeID, err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, networks)
	return
}

func (api *Api) CreateNodeNetwork(ctx *gin.Context) {
	var netWorkOption goclient.CreateNetworkOptions
	rolexContext, _ := ctx.Get("rolexContext")

	if err := ctx.BindJSON(&netWorkOption); err != nil {
		log.Error("create node network request body parse json error: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateNetworkParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	network, err := api.GetDockerClient().CreateNodeNetwork(rolexContext.(context.Context), netWorkOption)
	if err != nil {
		log.Error("create network error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, network)
	return
}
