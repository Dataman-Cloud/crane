package api

import (
	"net/http"
	"strings"

	"github.com/Dataman-Cloud/rolex/version"
	"github.com/gin-gonic/gin"
)

type RolexConfigResponse struct {
	Version      string `json:"version"`
	BuildTime    string `json:"build"`
	FeatureFlags string `json:"feature_flags"`
}

func (api *Api) RolexConfig(ctx *gin.Context) {
	config := &RolexConfigResponse{}
	config.Version = version.Version
	config.BuildTime = version.BuildTime
	config.FeatureFlags = strings.Join(api.GetConfig().FeatureFlags, ",")

	ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": config})
}

func (api *Api) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": "ok"})
}

func (api *Api) Help(engine *gin.Engine) gin.HandlerFunc {
	type RouteInfo struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}
	routes := make([]*RouteInfo, 0)
	for _, r := range engine.Routes() {
		routes = append(routes, &RouteInfo{
			Method: r.Method,
			Path:   r.Path,
		})
	}

	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "data": routes})
	}
}
