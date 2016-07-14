package api

import (
	"net/http"

	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectNode(ctx *gin.Context) {
	if ctx.Param("id") == "leader_manager" { // ugly will remove later
		api.LeaderNode(ctx)
	} else {
		node, err := api.GetDockerClient().InspectNode(ctx.Param("id"))
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
func (api *Api) UpdateNode(ctx *gin.Context) {}

func (api *Api) RemoveNode(ctx *gin.Context) {
	err := api.GetDockerClient().RemoveNode(ctx.Param("id"))
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

	ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": ""})
}

func (api *Api) Info(ctx *gin.Context) {
	info, err := api.GetDockerClient().Info(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": info})
}
