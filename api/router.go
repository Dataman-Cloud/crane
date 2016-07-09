package api

import (
	"github.com/Dataman-Cloud/rolex/api/middlewares"
	"github.com/Dataman-Cloud/rolex/util/log"
	"github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"time"
)

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()
	router.Use(log.Ginrus(logrus.StandardLogger(), time.RFC3339, true), gin.Recovery())

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
