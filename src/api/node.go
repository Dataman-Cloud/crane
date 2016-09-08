package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/crane/src/model"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/dmgin"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeUpdateNodeParamError = "400-11301"
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
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, node)
	return
}

func (api *Api) ManagerInfo(ctx *gin.Context) {
	systemInfo, err := api.GetDockerClient().ManagerInfo()
	if err != nil {
		log.Error("LeaderNode got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	nodeId := systemInfo.Swarm.NodeID
	node, err := api.GetDockerClient().InspectNode(nodeId)
	if err != nil {
		log.Errorf("InspectNode of %s got error: %s", nodeId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, node)
	return
}

func (api *Api) ListNodes(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		log.Error("ListNode got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, nodes)
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
		rerror := cranerror.NewError(CodeUpdateNodeParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rerror)
		return
	}

	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().UpdateNode(nodeId, nodeUpdate); err != nil {
		log.Errorf("Update node %s got error: %s", nodeId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) RemoveNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().RemoveNode(nodeId); err != nil {
		log.Errorf("Remove node %s got error: %s", nodeId, err.Error())
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) Info(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	info, err := api.GetDockerClient().Info(craneContext.(context.Context))
	if err != nil {
		log.Error("Get node info got error: ", err)
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	dmgin.HttpOkResponse(ctx, info)
	return
}
