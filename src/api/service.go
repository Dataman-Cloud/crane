package api

import (
	"encoding/base64"
	"strings"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	CodeUpdateServiceParamError = "400-11401"
	CodeCreateServiceParamError = "400-11402"
	CodeScaleServiceParamError  = "400-11403"

	CodeListTaskParamError = "400-11404"
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
	httpresponse.Ok(ctx, encryptServiceId(ctx.Param("service_id")))
}

func (api *Api) UpdateServiceImage(ctx *gin.Context) {
	encryptedServicId := ctx.Param("service_id")
	serviceId, err := decryptServiceId(encryptedServicId)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(serviceId)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	service.Spec.TaskTemplate.ContainerSpec.Image = ctx.Query("image")

	if err := api.GetDockerClient().UpdateServiceAutoOption(service.ID, service.Version, service.Spec); err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

// update single service
// notice update service api only update service spec
// if network does not exist return error and don't add default label about stack and registry auth
func (api *Api) UpdateService(ctx *gin.Context) {
	var craneServiceSpec model.CraneServiceSpec

	if err := ctx.BindJSON(&craneServiceSpec); err != nil {
		craneError := cranerror.NewError(CodeUpdateServiceParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	serviceId := ctx.Param("service_id")
	if err := dockerclient.ValidateCraneServiceSpec(&craneServiceSpec); err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	if err := api.GetDockerClient().CheckServicePortConflicts(&craneServiceSpec, serviceId); err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	service, err := api.GetDockerClient().InspectServiceWithRaw(serviceId)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	netAttachConfigs := []swarm.NetworkAttachmentConfig{}
	var serviceAlias string
	serviceAlias = craneServiceSpec.Name
	splitServiceNames := strings.Split(craneServiceSpec.Name, "_")
	if len(splitServiceNames) == 2 {
		serviceAlias = splitServiceNames[1]
	}

	for _, network := range craneServiceSpec.Networks {
		netAttachConfigs = append(netAttachConfigs, swarm.NetworkAttachmentConfig{
			Target:  network,
			Aliases: []string{serviceAlias},
		})

	}

	swarmServiceSpec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name:   craneServiceSpec.Name,
			Labels: craneServiceSpec.Labels,
		},
		Mode:         craneServiceSpec.Mode,
		TaskTemplate: craneServiceSpec.TaskTemplate,
		EndpointSpec: craneServiceSpec.EndpointSpec,
		Networks:     netAttachConfigs,
		UpdateConfig: craneServiceSpec.UpdateConfig,
	}
	updateOpts := types.ServiceUpdateOptions{}
	if craneServiceSpec.RegistryAuth != "" {
		registryAuth, err := dockerclient.EncodedRegistryAuth(craneServiceSpec.RegistryAuth)
		if err != nil {
			httpresponse.Error(ctx, err)
			return
		}
		updateOpts.EncodedRegistryAuth = registryAuth

		if swarmServiceSpec.Labels == nil {
			swarmServiceSpec.Labels = make(map[string]string)
		}

		swarmServiceSpec.Annotations.Labels[dockerclient.LabelRegistryAuth] = craneServiceSpec.RegistryAuth
	} else {
		delete(swarmServiceSpec.Annotations.Labels, dockerclient.LabelRegistryAuth)
	}

	if err := api.GetDockerClient().UpdateService(service.ID, service.Version, swarmServiceSpec, updateOpts); err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) InspectService(ctx *gin.Context) {
	service, err := api.GetDockerClient().InspectServiceWithRaw(ctx.Param("service_id"))
	if err != nil {
		log.Errorf("inspect service error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	craneService := model.CraneService{
		ID:           service.ID,
		Meta:         service.Meta,
		Spec:         api.GetDockerClient().ToCraneServiceSpec(service.Spec),
		Endpoint:     service.Endpoint,
		UpdateStatus: service.UpdateStatus,
	}
	httpresponse.Ok(ctx, craneService)
	return
}

// ServiceCreate creates a new Service.
func (api *Api) CreateService(ctx *gin.Context) {
	var service swarm.ServiceSpec
	if err := ctx.BindJSON(&service); err != nil {
		log.Error("CreateService invalied request body: ", err)
		craneError := cranerror.NewError(CodeCreateServiceParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	response, err := api.GetDockerClient().CreateService(service, types.ServiceCreateOptions{})
	if err != nil {
		log.Error("CreateService got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, response)
	return
}

func (api *Api) RemoveService(ctx *gin.Context) {
	serviceId := ctx.Param("id")
	if err := api.GetDockerClient().RemoveService(serviceId); err != nil {
		log.Errorf("Remove service %s got error: %s", serviceId, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) ScaleService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")
	var serviceScale dockerclient.ServiceScale
	if err := ctx.BindJSON(&serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		craneError := cranerror.NewError(CodeScaleServiceParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	if err := api.GetDockerClient().ScaleService(serviceId, serviceScale); err != nil {
		log.Errorf("Scale service %s got error: %s", serviceId, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
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
		craneError := cranerror.NewError(CodeListTaskParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	for _, task := range tasks {
		logContext := context.WithValue(context.Background(), "node_id", task.NodeID)
		go api.GetDockerClient().LogsContainer(logContext, task.Status.ContainerStatus.ContainerID, message)
	}

	w := ctx.Writer
	clientGone := w.CloseNotify()
	for {
		select {
		case data := <-message:
			ctx.SSEvent(dockerclient.SseTypeServiceLogs, data)
			w.Flush()
		case <-clientGone:
			return
		}
	}
}

func (api *Api) StatsService(ctx *gin.Context) {
	serviceId := ctx.Param("service_id")
	taskFilter := filters.NewArgs()
	taskFilter.Add("service", serviceId)

	tasks, err := api.GetDockerClient().ListTasks(types.TaskListOptions{Filter: taskFilter})
	if err != nil {
		log.Errorf("ListTasks of service %s got error: %s", serviceId, err.Error())
		craneError := cranerror.NewError(CodeListTaskParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	chnMsg := make(chan *model.CraneContainerStat, 1)
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
		opts.CraneContainerStats = chnMsg

		statsOptionsMap[opts.ID] = *opts

		go func(ctx context.Context, opts model.ContainerStatOptions) {
			chnErr <- api.GetDockerClient().StatsContainer(ctx, opts)
		}(statsContext, *opts)
	}

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
				ctx.SSEvent(dockerclient.SseTypeServiceStats, data)
				w.Flush()
			}
		case err := <-chnErr:
			if statsStopErr, ok := err.(*cranerror.ContainerStatsStopError); ok {
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
