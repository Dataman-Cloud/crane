package catalog

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Dataman-Cloud/rolex/src/util/config"

	"github.com/gin-gonic/gin"
)

type CatalogApi struct {
	Config *config.Config
}

func (catalogApi *CatalogApi) GetCatalog(ctx *gin.Context) {
	catalog, err := CatalogFromPath(filepath.Join(catalogApi.Config.CatalogPath, ctx.Param("name")))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": catalog})
}

func (catalogApi *CatalogApi) ListCatalog(ctx *gin.Context) {
	catalogs, err := AllCatalogFromPath(catalogApi.Config.CatalogPath)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": catalogs})
}
