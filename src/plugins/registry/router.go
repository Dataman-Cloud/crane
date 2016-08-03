package registry

import (
	"github.com/gin-gonic/gin"
)

func (registry *Registry) RegisterApiForRegistry(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	registryV1 := router.Group("/registry/v1")
	{
		registryV1.GET("/token", registry.Token)
		registryV1.POST("/notifications", registry.Notifications)
	}

	registryV1Protected := router.Group("/registry/v1", middlewares...)
	{
		registryV1Protected.GET("/catalogs/mine", registry.MineCatalog)
		registryV1Protected.GET("/catalogs/public", registry.PublicCatalog) // under library or tag marked as public
		registryV1Protected.GET("/tag/list/:namespace/:image", registry.TagList)
		registryV1Protected.GET("/manifests/:reference/:namespace/:image", registry.GetManifests)
		registryV1Protected.PATCH("/:namespace/:image/publicity", registry.ImagePublicity)
		registryV1Protected.DELETE("/manifests/:reference/:namespace/:image", registry.DeleteManifests)
	}
}
