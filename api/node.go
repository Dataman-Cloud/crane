package api

import (
	"encoding/json"
	"net/http"

	"github.com/Dataman-Cloud/rolex/util"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Node2Update represents a node spec to update.
type Node2Update struct {
	Spec    swarm.NodeSpec
	Version swarm.Version
}

func (api *Api) InspectNode(ctx *gin.Context) {
	if ctx.Param("node_id") == "leader_manager" { // ugly will remove later
		api.LeaderNode(ctx)
	} else {
		node, err := api.GetDockerClient().InspectNode(ctx.Param("node_id"))
		if err != nil {
			ctx.JSON(http.StatusServiceUnavailable, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": node})
	}
}

func (api *Api) ListNodes(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": nodes})
}

func (api *Api) CreateNode(ctx *gin.Context) {}

func (api *Api) UpdateNode(ctx *gin.Context) {
	var node2Update Node2Update

	if err := ctx.BindJSON(&node2Update); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("Node JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}

		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := api.GetDockerClient().UpdateNode(ctx.Param("node_id"), node2Update.Version, node2Update.Spec)
	if err != nil {
		log.Errorf("UpdateNode: spec error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS})
}

func (api *Api) RemoveNode(ctx *gin.Context) {
	err := api.GetDockerClient().RemoveNode(ctx.Param("node_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "0"})
}

func (api *Api) LeaderNode(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	for _, node := range nodes {
		if node.ManagerStatus != nil && node.ManagerStatus.Leader {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": node.ManagerStatus})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": ""})
}

func (api *Api) Info(ctx *gin.Context) {
	rolexContext, _ := ctx.Get("rolexContext")
	info, err := api.GetDockerClient().Info(rolexContext.(context.Context))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": info})
}
