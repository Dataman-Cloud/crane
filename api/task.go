package api

import (
	"net/http"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectTask(ctx *gin.Context) {
	task, err := api.GetDockerClient().InspectTask(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": task})
}

func (api *Api) ListTasks(ctx *gin.Context) {
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", ctx.Param("service_id"))

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tasks})
}
