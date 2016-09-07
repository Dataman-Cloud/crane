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
	CodeCatalogGetCatalogError    = "503-15031"
	CodeCatalogListCatalogError   = "503-15032"
	CodeCatalogInvalidUser        = "401-15033"
	CodeCatalogInvalidCatalogId   = "403-15034"
	CodeCatalogInvalidParam       = "400-15035"
	CodeCatalogInvalidIcon        = "400-15036"
	CodeCatalogDeleteFaild        = "503-15037"
	CodeCatalogForbiddenOperation = "403-15038"
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

func (catalogApi *CatalogApi) CreateCatalog(ctx *gin.Context) {
	var catalog Catalog
	if err := ctx.BindJSON(&catalog); err != nil {
		log.Errorf("invalid param error: %v", err)
		rolexerr := dmerror.NewError(CodeCatalogInvalidParam, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if iconData, err := ImageHandle(ctx.Request); err != nil {
		rolexerr := dmerror.NewError(CodeCatalogInvalidIcon, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		catalog.IconData = iconData
	}

	catalogApi.Save(&catalog)
	dmgin.HttpOkResponse(ctx, catalog)
}

func (catalogApi *CatalogApi) DeleteCatalog(ctx *gin.Context) {
	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		log.Error("invalid catalog_id")
		rolexerr := dmerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err = catalogApi.Delete(catalogId); err != nil {
		log.Errorf("delete catalog error: %v", err)
		rolexerr := dmerror.NewError(CodeCatalogDeleteFaild, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, "delete success")
}

func (catalogApi *CatalogApi) UpdateCatalog(ctx *gin.Context) {
	var catalog Catalog
	if err := ctx.BindJSON(&catalog); err != nil {
		log.Errorf("invalid param error: %v", err)
		rolexerr := dmerror.NewError(CodeCatalogInvalidParam, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		log.Error("invalid catalog_id")
		rolexerr := dmerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	cl, err := catalogApi.Get(catalogId)
	if err != nil {
		rolexerr := dmerror.NewError(CodeCatalogInvalidParam, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if cl.Type == CATALOG_SYSTEM_DEFAULT {
		rolexerr := dmerror.NewError(CodeCatalogForbiddenOperation, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if iconData, err := ImageHandle(ctx.Request); err != nil {
		rolexerr := dmerror.NewError(CodeCatalogInvalidIcon, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else if iconData == "" {
		catalog.IconData = cl.IconData
	} else {
		catalog.IconData = iconData
	}

	catalog.ID = cl.ID
	catalog.Name = cl.Name
	catalog.Bundle = cl.Bundle
	catalog.Readme = cl.Readme
	catalog.Description = cl.Description
	catalog.UserId = cl.UserId
	catalog.Type = cl.Type

	dmgin.HttpOkResponse(ctx, "update success")
}
