package catalog

import (
	"github.com/gin-gonic/gin"
)

func (catalogApi *CatalogApi) RegisterApiForCatalog(router *gin.Engine, middlewares ...gin.HandlerFunc) {

	catalogV1Protected := router.Group("/catalog/v1", middlewares...)
	{
		catalogV1Protected.GET("/catalogs", catalogApi.ListCatalog)
		catalogV1Protected.GET("/catalogs/:name", catalogApi.GetCatalog)
	}
}
