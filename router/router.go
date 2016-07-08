package router

import (
	"github.com/Dataman-Cloud/newworld/rolex-go/api"
	"github.com/Dataman-Cloud/newworld/rolex-go/router/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	groupv1 := router.Group("/api/v1", middlewares.Authorization)
	{
		router.GET("/health", api.HealthCheck)
	}

	return router
}
