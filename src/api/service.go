package api

import (
	"fmt"
	"io"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

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
		log.Error("UpdateService invalied request body: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeUpdateServiceParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	if err := api.GetDockerClient().UpdateService(service.ID, service.Version, serviceSpec, nil); err != nil {
		log.Errorf("update service error: %v", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) InspectService(ctx *gin.Context) {
	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, service)
	return
}

// ServiceCreate creates a new Service.
func (api *Api) CreateService(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		log.Error("CreateService invalied request body: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateServiceParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	response, err := api.GetDockerClient().CreateService(service, types.ServiceCreateOptions{})
	if err != nil {
		log.Error("CreateService got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, response)
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
		log.Error("ListService got error: ", err)
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, services)
	return
}

func (api *Api) RemoveService(ctx *gin.Context) {
	serviceId := ctx.Param("id")
	if err := api.GetDockerClient().RemoveService(serviceId); err != nil {
		log.Errorf("Remove service %s got error: %s", serviceId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) ScaleService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")
	var serviceScale dockerclient.ServiceScale
	if err := ctx.BindJSON(&serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeScaleServiceParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
		return
	}

	if err := api.GetDockerClient().ScaleService(serviceId, serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		api.HttpErrorResponse(ctx, err)
		return
	}

	api.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) LogsService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", serviceId)
	message := make(chan string)
	defer close(message)

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		log.Errorf("ListTasks of service %s got error: %s", serviceId, err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeListTaskParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
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
	serviceId := ctx.Param("service_id")
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", serviceId)
	stats := make(chan *model.ContainerStat)

	defer close(stats)

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		log.Errorf("ListTasks of service %s got error: %s", serviceId, err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeListTaskParamError, err.Error())
		api.HttpErrorResponse(ctx, rerror)
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
