package api

import (
	"github.com/Dataman-Cloud/rolex/api/middlewares"

	"github.com/gin-gonic/gin"
)

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	v1 := router.Group("/api/v1", middlewares.Authorization)
	{
		v1.GET("/health", api.HealthCheck)
		v1.GET("/nodes", api.ListNodes)
		v1.GET("/nodes/:id", api.InspectNode)
	}

	return router
}
