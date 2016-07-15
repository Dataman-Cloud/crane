package api

import (
	"time"

	"github.com/Dataman-Cloud/rolex/api/middlewares"
	"github.com/Dataman-Cloud/rolex/plugins/registry"
	"github.com/Dataman-Cloud/rolex/util/log"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()

	router.Use(log.Ginrus(logrus.StandardLogger(), time.RFC3339, true), gin.Recovery())
	router.Use(middlewares.OptionHandler())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	if api.Config.FeatureEnabled("registry") {
		r := &registry.Registry{Config: api.Config}
		r.RegisterApiForRegistry(router)
	}

	v1 := router.Group("/api/v1", middlewares.Authorization)
	{
		v1.GET("/nodes", api.ListNodes)
		v1.GET("/nodes/:id", api.InspectNode)
		v1.GET("/nodes/:id/info", api.Info)

		// Going to delegate to /nodes/:id
		// v1.GET("/nodes/leader_manager", api.LeaderNode)

		v1.GET("/nodes/:id/containers", api.ListContainers)
		v1.GET("/nodes/:id/containers/:container_id", api.InspectContainer)
		v1.GET("/nodes/:id/containers/:container_id/diff", api.DiffContainer)
		v1.DELETE("/nodes/:id/containers/:container_id", api.RemoveContainer)
		v1.DELETE("/nodes/:id/containers/:container_id/kill", api.KillContainer)

		v1.POST("/networks", api.CreateNetwork)
		v1.GET("/networks", api.ListNetworks)
		v1.DELETE("/networks/:id", api.RemoveNetwork)
		v1.GET("/networks/:id", api.InspectNetwork)
		v1.PATCH("/networks/:id", api.ConnectNetwork)

		v1.POST("/stacks", api.CreateStack)
		v1.GET("/stacks", api.ListStack)
		v1.GET("/stacks/:namespace", api.InspectStack)
		v1.DELETE("/stacks/:namespace", api.RemoveStack)
		v1.PATCH("/stacks/:namespace/services/:serviceID", api.ScaleService)
		v1.GET("/stacks/:namespace/services", api.ListStackService)

		v1.DELETE("/volumes/:node_id/:name", api.RemoveVolume)
		v1.GET("/volumes/:node_id/:name", api.InspectVolume)
		v1.GET("/volumes/:node_id", api.ListVolume)
		v1.POST("/volumes/:node_id", api.CreateVolume)

		v1.GET("/images/:node_id", api.ListImages)
		v1.GET("/images/:node_id/:name", api.InspectImage)
		v1.GET("/images/:node_id/:name/history", api.ImageHistory)
	}

	misc := router.Group("/misc/v1")
	{
		misc.GET("/help", api.Help(router))
		misc.GET("/config", api.RolexConfig)
		misc.GET("/health", api.HealthCheck)
	}

	return router
}
