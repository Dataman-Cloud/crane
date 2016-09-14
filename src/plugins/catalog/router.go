package catalog

import (
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/utils/db"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Init() {
	log.Infof("begin to init and enable plugin: %s", apiplugin.Catalog)
	catalogApi := &CatalogApi{DbClient: db.DB()}
	catalogApi.DbClient.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&Catalog{})

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.Catalog,
		Dependencies: []string{apiplugin.Db},
		Instance:     catalogApi,
	}

	apiplugin.Add(apiPlugin)
	log.Infof("init and enable plugin: %s success", apiplugin.Catalog)
}

func (catalogApi *CatalogApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	catalogV1 := router.Group("/catalog/v1", middlewares...)
	{
		catalogV1.GET("/catalogs", catalogApi.ListCatalog)
		catalogV1.POST("/catalogs", catalogApi.CreateCatalog)

		catalogV1.GET("/catalogs/:catalog_id", catalogApi.GetCatalog)
		catalogV1.PATCH("/catalogs/:catalog_id", catalogApi.UpdateCatalog)
		catalogV1.DELETE("/catalogs/:catalog_id", catalogApi.DeleteCatalog)
	}
}
