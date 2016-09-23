package api

import (
	"net/http"
	"runtime"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"
	"github.com/Dataman-Cloud/crane/src/version"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

const (
	//Get config error code
	CodeGetConfigError = "503-11901"
)

type CraneConfigResponse struct {
	Version      string      `json:"Version"`
	BuildTime    string      `json:"Build"`
	FeatureFlags []string    `json:"FeatureFlags"`
	SwarmInfo    swarm.Swarm `json:"SwarmInfo"`
	NumGoroutine int
}

type RouteInfo struct {
	Method string `json:"Method"`
	Path   string `json:"Path"`
}

func (api *Api) CraneConfig(ctx *gin.Context) {
	config := &CraneConfigResponse{}
	config.Version = version.Version
	config.BuildTime = version.BuildTime
	config.FeatureFlags = api.GetConfig().FeatureFlags
	config.NumGoroutine = runtime.NumGoroutine()

	var err error
	config.SwarmInfo, err = api.GetDockerClient().InspectSwarm()

	if err != nil {
		log.Errorf("InspectSwarm got error: %s", err.Error())
		craneError := cranerror.NewError(CodeGetConfigError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	httpresponse.Ok(ctx, config)
	return
}

func (api *Api) HealthCheck(ctx *gin.Context) {
	// node docker client check
	nodes, err := api.GetDockerClient().ListNode(types.NodeListOptions{})
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	var craneContext context.Context
	backgroundContext := context.Background()

	for _, node := range nodes {
		if node.Status.State != swarm.NodeStateReady {
			continue
		}

		craneContext = context.WithValue(backgroundContext, "node_id", node.ID)
		_, err = api.GetDockerClient().SwarmNode(craneContext)
		if err != nil {
			httpresponse.Error(ctx, err)
			return
		}
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) Help(engine *gin.Engine) gin.HandlerFunc {
	routes := make([]*RouteInfo, 0)
	for _, r := range engine.Routes() {
		routes = append(routes, &RouteInfo{
			Method: r.Method,
			Path:   r.Path,
		})
	}

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": routes})
	}
}
