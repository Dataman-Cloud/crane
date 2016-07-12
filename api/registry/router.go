package registry

import (
	"github.com/gin-gonic/gin"
)

func (registry *Registry) RegisterApiForRegistry(router *gin.Engine) {
	registryV1 := router.Group("/registry/v1")
	{
		registryV1.GET("/token", registry.Token)
	}
}
