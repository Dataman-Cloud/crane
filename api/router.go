package api

import (
	"time"

	"github.com/Dataman-Cloud/rolex/api/middlewares"
	"github.com/Dataman-Cloud/rolex/util/log"
	"github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()
	router.Use(log.Ginrus(logrus.StandardLogger(), time.RFC3339, true), gin.Recovery())

	//router.Use(auth.OptionHandler)
	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	v1 := router.Group("/api/v1", middlewares.Authorization)
	{
		v1.GET("/health", api.HealthCheck)
		v1.GET("/nodes", api.ListNodes)
		v1.GET("/nodes/:id", api.InspectNode)

		v1.GET("/containers", api.ListContainers)
		v1.GET("/containers/:id", api.InspectContainer)

		v1.POST("/services/create", api.ServiceCreate)
		v1.GET("/services", api.ServiceList)
		v1.DELETE("/services/:id", api.ServiceRemove)

		v1.POST("/networks/:id/container", api.ConnectNetwork)
		v1.DELETE("/networks/:id/container", api.DisconnectNetwork)

		v1.POST("/networks", api.CreateNetwork)
		v1.GET("/networks/:id", api.InspectNetwork)
		v1.GET("/networks", api.ListNetworks)
		v1.DELETE("/networks/:id", api.RemoveNetwork)
	}

	return router
}

func OptionHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, X-Requested-By, If-Modified-Since, X-File-Name, Cache-Control, X-XSRFToken, Authorization")
	ctx.Header("Content-Type", "application/json")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
	}

	ctx.Next()
}
