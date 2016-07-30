package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/manucorporat/sse"
	"golang.org/x/net/context"
)

func (api *Api) UpdateService(ctx *gin.Context) {
	var serviceSpec swarm.ServiceSpec

	if err := ctx.BindJSON(&serviceSpec); err != nil {
		log.Errorf("invalied request body: %v", err)
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	if err := api.GetDockerClient().UpdateService(service.ID, service.Version, serviceSpec, nil); err != nil {
		log.Errorf("update service error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": "update success"})
}

func (api *Api) InspectService(ctx *gin.Context) {
	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": util.ENGINE_OPERATION_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": util.OPERATION_SUCCESS, "data": service})
}

// ServiceCreate creates a new Service.
func (api *Api) CreateService(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	response, err := api.GetDockerClient().CreateService(service, types.ServiceCreateOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": response})
	return
}

// ServiceList returns the list of services.
func (api *Api) ListService(ctx *gin.Context) {
	opts := types.ServiceListOptions{}
	if labelFilters_, found := ctx.Get("labelFilters"); found {
		labelFilters := labelFilters_.(map[string]string)
		args := filters.NewArgs()
		for k, v := range labelFilters {
			args.Add("label", fmt.Sprintf("%s=%s", k, v))
		}
		fmt.Println(args)
		opts.Filter = args
	}

	services, err := api.GetDockerClient().ListService(opts)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": services})
	return
}

func (api *Api) RemoveService(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	if serviceID == "" {
		ctx.JSON(http.StatusBadRequest, "service id is nil")
		return
	}

	if err := api.GetDockerClient().RemoveService(serviceID); err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
	return
}

func (api *Api) ScaleService(ctx *gin.Context) {
	serviceID := ctx.Param("service_id")
	var serviceScale dockerclient.ServiceScale
	if err := ctx.BindJSON(&serviceScale); err != nil {
		log.Error("Scale service got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	if err := api.GetDockerClient().ScaleService(serviceID, serviceScale); err != nil {
		log.Error("Scale service got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
	return
}

func (api *Api) LogsService(ctx *gin.Context) {
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", ctx.Param("service_id"))
	message := make(chan string)
	defer close(message)

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	for _, task := range tasks {
		logContext := context.WithValue(context.Background(), "node_id", task.NodeID)
		go api.GetDockerClient().LogsContainer(logContext, task.Status.ContainerStatus.ContainerID, message)
	}

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "service-logs",
			Data:  <-message,
		}.Render(ctx.Writer)
		return true
	})
}

func (api *Api) StatsService(ctx *gin.Context) {
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", ctx.Param("service_id"))
	stats := make(chan *model.ContainerStat)

	defer close(stats)

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	for _, task := range tasks {
		logContext := context.WithValue(context.Background(), "node_id", task.NodeID)
		go api.GetDockerClient().StatsContainer(logContext, task.Status.ContainerStatus.ContainerID, stats)
	}

	ctx.Stream(func(w io.Writer) bool {
		sse.Event{
			Event: "service-stats",
			Data:  <-stats,
		}.Render(ctx.Writer)
		return true
	})
}
