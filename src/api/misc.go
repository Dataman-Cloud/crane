package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/version"
	"github.com/gin-gonic/gin"
)

type RolexConfigResponse struct {
	Version      string   `json:"Version"`
	BuildTime    string   `json:"Build"`
	FeatureFlags []string `json:"FeatureFlags"`
	RolexSecret  string   `json:"RolexSecret"`
	RolexCaHash  string   `json:"RolexCaHash"`
}

func (api *Api) RolexConfig(ctx *gin.Context) {
	config := &RolexConfigResponse{}
	config.Version = version.Version
	config.BuildTime = version.BuildTime
	config.FeatureFlags = api.GetConfig().FeatureFlags

	config.RolexSecret = api.GetConfig().RolexSecret
	config.RolexCaHash = api.GetConfig().RolexCaHash
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": config})
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
