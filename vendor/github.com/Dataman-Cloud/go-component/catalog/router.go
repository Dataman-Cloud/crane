package catalog

import (
	"github.com/Dataman-Cloud/rolex/src/utils/db"

	"github.com/gin-gonic/gin"
)

func NewCatalog(catalogPath string) *CatalogApi {
	return &CatalogApi{
		CatalogPath: catalogPath,
		DbClient:    db.DB(),
	}
}

func (catalogApi *CatalogApi) MigriateTable() {
	catalogApi.DbClient.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&Catalog{})
}

func (catalogApi *CatalogApi) RegisterApiForCatalog(router *gin.Engine, middlewares ...gin.HandlerFunc) {

	catalogV1 := router.Group("/catalog/v1", middlewares...)
	{
		catalogV1.GET("/catalogs", catalogApi.ListCatalog)
		catalogV1.POST("/catalogs", catalogApi.CreateCatalog)

		catalogV1.GET("/catalogs/:catalog_id", catalogApi.GetCatalog)
		catalogV1.PATCH("/catalogs/:catalog_id", catalogApi.UpdateCatalog)
		catalogV1.DELETE("/catalogs/:catalog_id", catalogApi.DeleteCatalog)
	}
}
