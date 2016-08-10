package api

import (
	"encoding/base64"
	"fmt"
	"io"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"github.com/manucorporat/sse"
	"golang.org/x/net/context"
)

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func decryptServiceId(encryptServiceId string) (string, error) {
	serviceId, err := base64.StdEncoding.DecodeString(reverseString(encryptServiceId))
	if err != nil {
		return "", err
	}

	return string(serviceId), nil
}

func encryptServiceId(serviceId string) string {
	return reverseString(base64.StdEncoding.EncodeToString([]byte(serviceId)))
}

func (api *Api) ServiceCDAddr(ctx *gin.Context) {
	rolexgin.HttpOkResponse(ctx, encryptServiceId(ctx.Param("service_id")))
}

func (api *Api) UpdateServiceImage(ctx *gin.Context) {
	encryptedServicId := ctx.Param("service_id")
	serviceId, err := decryptServiceId(encryptedServicId)
	if err != nil {
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(serviceId)
	if err != nil {
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	service.Spec.TaskTemplate.ContainerSpec.Image = ctx.Query("image")

	if err := api.GetDockerClient().UpdateService(service.ID, service.Version, service.Spec, nil); err != nil {
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) UpdateService(ctx *gin.Context) {
	var serviceSpec swarm.ServiceSpec

	if err := ctx.BindJSON(&serviceSpec); err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodeUpdateServiceParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	if err := api.GetDockerClient().UpdateService(service.ID, service.Version, serviceSpec, nil); err != nil {
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) InspectService(ctx *gin.Context) {
	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, service)
	return
}

// ServiceCreate creates a new Service.
func (api *Api) CreateService(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		log.Error("CreateService invalied request body: ", err)
		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateServiceParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	response, err := api.GetDockerClient().CreateService(service, types.ServiceCreateOptions{})
	if err != nil {
		log.Error("CreateService got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, response)
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
		opts.Filter = args
	}

	services, err := api.GetDockerClient().ListService(opts)
	if err != nil {
		log.Error("ListService got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, services)
	return
}

func (api *Api) RemoveService(ctx *gin.Context) {
	serviceId := ctx.Param("id")
	if err := api.GetDockerClient().RemoveService(serviceId); err != nil {
		log.Errorf("Remove service %s got error: %s", serviceId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) ScaleService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")
	var serviceScale dockerclient.ServiceScale
	if err := ctx.BindJSON(&serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeScaleServiceParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	if err := api.GetDockerClient().ScaleService(serviceId, serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
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
		rolexgin.HttpErrorResponse(ctx, rerror)
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

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		log.Errorf("ListTasks of service %s got error: %s", serviceId, err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeListTaskParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	chnMsg := make(chan *model.ContainerStat, 1)
	defer close(chnMsg)
	chnErr := make(chan error, 1)
	defer close(chnErr)

	statsOptionsMap := make(map[string]model.ContainerStatOptions)
	for _, task := range tasks {
		if task.Status.State != swarm.TaskStateRunning {
			continue
		}

		statsContext := context.WithValue(context.Background(), "node_id", task.NodeID)
		opts := createStatOption()
		opts.ID = task.Status.ContainerStatus.ContainerID
		opts.RolexContainerStats = chnMsg

		statsOptionsMap[opts.ID] = *opts

		go func(ctx context.Context, opts model.ContainerStatOptions) {
			chnErr <- api.GetDockerClient().StatsContainer(ctx, opts)
		}(statsContext, *opts)
	}

	ssEvent := &sse.Event{Event: "service-stats"}
	w := ctx.Writer
	clientGone := w.CloseNotify()
	clientClosed := false

	for {
		select {
		case <-clientGone:
			clientClosed = true
			log.Infof("Stats stream of service %s closed by client", serviceId)
			for _, statOpts := range statsOptionsMap {
				statOpts.Done <- true
			}
		case data := <-chnMsg:
			if !clientClosed {
				ssEvent.Data = data
				ssEvent.Render(w)
			}
		case err := <-chnErr:
			if statsStopErr, ok := err.(*rolexerror.ContainerStatsStopError); ok {
				containerId := statsStopErr.ID
				log.Infof("Stats stream of container %s stop with error: %s", containerId, statsStopErr.Error())
				if statOpts, ok := statsOptionsMap[containerId]; ok {
					close(statOpts.Done)
					delete(statsOptionsMap, containerId)
					if len(statsOptionsMap) == 0 {
						log.Infof("Stats stream of service %s stop", serviceId)
						return
					}
				}
			} else {
				log.Error("Received unknown error: ", err)
			}
		}
	}
}

func createStatOption() *model.ContainerStatOptions {
	chnDone := make(chan bool, 1)                    //chosed by func StatsService
	chnContainerStats := make(chan *docker.Stats, 1) // closed by go-dockerclient
	return &model.ContainerStatOptions{
		Stats:  chnContainerStats,
		Stream: true,
		Done:   chnDone,
	}

}
