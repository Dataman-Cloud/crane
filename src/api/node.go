package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")

	if nodeId == "manager_info" { // ugly will remove later
		api.ManagerInfo(ctx)
		return
	}

	node, err := api.GetDockerClient().InspectNode(nodeId)
	if err != nil {
		log.Errorf("InspectNode of %s got error: %s", nodeId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, node)
	return
}

func (api *Api) ManagerInfo(ctx *gin.Context) {
	systemInfo, err := api.GetDockerClient().ManagerInfo()
	if err != nil {
		log.Error("LeaderNode got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	nodeId := systemInfo.Swarm.NodeID
	node, err := api.GetDockerClient().InspectNode(nodeId)
	if err != nil {
		log.Errorf("InspectNode of %s got error: %s", nodeId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, node)
	return
}

func (api *Api) ListNodes(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		log.Error("ListNode got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, nodes)
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
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().UpdateNode(nodeId, nodeUpdate); err != nil {
		log.Errorf("Update node %s got error: %s", nodeId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) RemoveNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().RemoveNode(nodeId); err != nil {
		log.Errorf("Remove node %s got error: %s", nodeId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) Info(ctx *gin.Context) {
	info, err := api.GetDockerClient().Info(ctx.Param("node_id"))
	if err != nil {
		log.Error("Get node info got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, info)
	return
}
