package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (api *Api) InspectNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")
	if nodeId == "leader_manager" { // ugly will remove later
		api.LeaderNode(ctx)
		return
	}
	node, err := api.GetDockerClient().InspectNode(nodeId)
	if err != nil {
		log.Errorf("InspectNode of %s got error: %s", nodeId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, node)
	return
}

func (api *Api) ListNodes(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		log.Error("ListNode got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, nodes)
	return
}

func (api *Api) CreateNode(ctx *gin.Context) {}

func (api *Api) UpdateNode(ctx *gin.Context) {
	var nodeUpdate model.UpdateOptions

	if err := ctx.BindJSON(&nodeUpdate); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("Node JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
		rerror := rolexerror.NewRolexError(rolexerror.CodeUpdateNodeParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().UpdateNode(nodeId, nodeUpdate); err != nil {
		log.Errorf("Update node %s got error: %s", nodeId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) RemoveNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().RemoveNode(nodeId); err != nil {
		log.Errorf("Remove node %s got error: %s", nodeId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) LeaderNode(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		log.Error("LeaderNode got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	for _, node := range nodes {
		if node.ManagerStatus != nil && node.ManagerStatus.Leader {
			api.HttpOkResponse(ctx, node.ManagerStatus)
			return
		}
	}

	api.HttpOkResponse(ctx, "")
	return
}

func (api *Api) Info(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	info, err := api.GetDockerClient().Info(rolexContext.(context.Context))
	if err != nil {
		log.Error("Get node info got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, info)
	return
}
