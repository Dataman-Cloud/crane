package catalog

import (
	"path/filepath"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/go-component/utils/dmgin"

	"github.com/gin-gonic/gin"
)

const (
	//Catalog
	CodeCatalogGetCatalogError  = "503-15031"
	CodeCatalogListCatalogError = "503-15032"
)

type CatalogApi struct {
	CatalogPath string
}

func (catalogApi *CatalogApi) GetCatalog(ctx *gin.Context) {
	catalog, err := CatalogFromPath(filepath.Join(catalogApi.CatalogPath, ctx.Param("name")))
	if err != nil {
		rolexerr := dmerror.NewError(CodeCatalogGetCatalogError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, catalog)
}

func (catalogApi *CatalogApi) ListCatalog(ctx *gin.Context) {
	catalogs, err := AllCatalogFromPath(catalogApi.CatalogPath)
	if err != nil {
		rolexerr := dmerror.NewError(CodeCatalogListCatalogError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, catalogs)
}
