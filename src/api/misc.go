package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"
	"github.com/Dataman-Cloud/rolex/src/version"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
)

type RolexConfigResponse struct {
	Version      string      `json:"Version"`
	BuildTime    string      `json:"Build"`
	FeatureFlags []string    `json:"FeatureFlags"`
	SwarmInfo    swarm.Swarm `json:"SwarmInfo"`
}

func (api *Api) RolexConfig(ctx *gin.Context) {
	config := &RolexConfigResponse{}
	config.Version = version.Version
	config.BuildTime = version.BuildTime
	config.FeatureFlags = api.GetConfig().FeatureFlags

	var err error
	config.SwarmInfo, err = api.GetDockerClient().InspectSwarm()

	if err != nil {
		log.Errorf("InspectSwarm got error: %s", err.Error())
		rerror := rolexerror.NewRolexError(rolexerror.CodeGetConfigError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	rolexgin.HttpOkResponse(ctx, config)
	return
}

func (api *Api) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "ok"})
}

func (api *Api) Help(engine *gin.Engine) gin.HandlerFunc {
	type RouteInfo struct {
		Method string `json:"Method"`
		Path   string `json:"Path"`
	}
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
