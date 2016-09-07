package catalog

import (
	"strconv"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/go-component/utils/dmgin"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

const (
	//Catalog
	CodeCatalogGetCatalogError  = "503-15031"
	CodeCatalogListCatalogError = "503-15032"
	CodeCatalogInvalidUser      = "401-15033"
	CodeCatalogInvalidCatalogId = "403-15034"
)

type CatalogApi struct {
	CatalogPath string
	DbClient    *gorm.DB
}

func (catalogApi *CatalogApi) GetCatalog(ctx *gin.Context) {
	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	catalog, err := catalogApi.Get(catalogId)
	if err != nil {
		log.Errorf("get catalog error: %v", err)
		rolexerr := dmerror.NewError(CodeCatalogGetCatalogError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, catalog)
}

func (catalogApi *CatalogApi) ListCatalog(ctx *gin.Context) {
	//account, _ := ctx.Get("account")
	catalogs, err := catalogApi.List()
	if err != nil {
		log.Errorf("get catalog list error: %v", err)
		rolexerr := dmerror.NewError(CodeCatalogInvalidUser, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, catalogs)
}
