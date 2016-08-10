package catalog

import (
	"net/http"
	"path/filepath"

	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	"github.com/gin-gonic/gin"
)

type CatalogApi struct {
	Config *config.Config
}

func (catalogApi *CatalogApi) GetCatalog(ctx *gin.Context) {
	catalog, err := CatalogFromPath(filepath.Join(catalogApi.Config.CatalogPath, ctx.Param("name")))
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeCatalogGetCatalogError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	rolexgin.HttpOkResponse(ctx, catalog)
}

func (catalogApi *CatalogApi) ListCatalog(ctx *gin.Context) {
	catalogs, err := AllCatalogFromPath(catalogApi.Config.CatalogPath)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeCatalogListCatalogError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	rolexgin.HttpOkResponse(ctx, catalogs)
}
