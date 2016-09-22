package api

import (
	"encoding/json"

	"github.com/Dataman-Cloud/crane/src/model"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeUpdateNodeParamError = "400-11301"
	CodeCreateNodeParamError = "400-11306"
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
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, node)
	return
}

func (api *Api) ManagerInfo(ctx *gin.Context) {
	systemInfo, err := api.GetDockerClient().ManagerInfo()
	if err != nil {
		log.Error("LeaderNode got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	nodeId := systemInfo.Swarm.NodeID
	node, err := api.GetDockerClient().InspectNode(nodeId)
	if err != nil {
		log.Errorf("InspectNode of %s got error: %s", nodeId, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, node)
	return
}

func (api *Api) ListNodes(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		log.Error("ListNode got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, nodes)
	return
}

func (api *Api) CreateNode(ctx *gin.Context) {
	var joiningNode model.JoiningNode

	if err := ctx.BindJSON(&joiningNode); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
		rerror := cranerror.NewError(CodeCreateNodeParamError, err.Error())
		httpresponse.Error(ctx, rerror)
		return
	}

	if err := api.GetDockerClient().CreateNode(joiningNode); err != nil {
		log.Errorf("Create node %s got error: %s", joiningNode.Endpoint, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

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
		craneError := cranerror.NewError(CodeUpdateNodeParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().UpdateNode(nodeId, nodeUpdate); err != nil {
		log.Errorf("Update node %s got error: %s", nodeId, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) RemoveNode(ctx *gin.Context) {
	nodeId := ctx.Param("node_id")
	if err := api.GetDockerClient().RemoveNode(nodeId); err != nil {
		log.Errorf("Remove node %s got error: %s", nodeId, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) Info(ctx *gin.Context) {
	craneContext, _ := ctx.Get("craneContext")
	info, err := api.GetDockerClient().Info(craneContext.(context.Context))
	if err != nil {
		log.Error("Get node info got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, info)
	return
}
