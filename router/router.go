package router

import (
	"github.com/Dataman-Cloud/rolex/api"
	"github.com/Dataman-Cloud/rolex/router/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	v1 := router.Group("/api/v1", middlewares.Authorization)
	{
		v1.GET("/health", api.HealthCheck)
		v1.GET("/nodes", api.GetNodes)
	}

	return router
}
