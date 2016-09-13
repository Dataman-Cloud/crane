package catalog

import (
	"strconv"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

const (
	//Catalog
	CodeCatalogGetCatalogError    = "404-15031"
	CodeCatalogListCatalogError   = "503-15032"
	CodeCatalogInvalidUser        = "401-15033"
	CodeCatalogInvalidCatalogId   = "400-15034"
	CodeCatalogInvalidParam       = "400-15035"
	CodeCatalogInvalidIcon        = "400-15036"
	CodeCatalogDeleteFaild        = "503-15037"
	CodeCatalogForbiddenOperation = "403-15038"
	CodeCatalogUpdateFaild        = "503-15039"
	CodeCatalogCreateFaild        = "503-15040"
)

const (
	CATALOG_SYSTEM_DEFAULT = 0
	CATALOG_USER_CUSTOM    = 1
)

type CatalogApi struct {
	CatalogPath string
	DbClient    *gorm.DB
}

func (catalogApi *CatalogApi) GetCatalog(ctx *gin.Context) {
	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	catalog, err := catalogApi.Get(catalogId)
	if err != nil {
		log.Errorf("get catalog error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, catalog)
}

func (catalogApi *CatalogApi) ListCatalog(ctx *gin.Context) {
	catalogs, err := catalogApi.List()
	if err != nil {
		log.Errorf("get catalog list error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, catalogs)
}

func (catalogApi *CatalogApi) CreateCatalog(ctx *gin.Context) {
	account, ok := ctx.Get("account")
	if !ok {
		craneerr := cranerror.NewError(CodeCatalogInvalidUser, "invalid user login")
		httpresponse.Error(ctx, craneerr)
		return
	}

	var catalog Catalog
	if err := ctx.Bind(&catalog); err != nil {
		log.Errorf("invalid param error: %v", err)
		craneerr := cranerror.NewError(CodeCatalogInvalidParam, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if iconData, err := ImageHandle(ctx.Request); err != nil {
		craneerr := cranerror.NewError(CodeCatalogInvalidIcon, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	} else {
		catalog.IconData = iconData
		catalog.Type = CATALOG_USER_CUSTOM
	}

	if catalog.Bundle == "" {
		craneerr := cranerror.NewError(CodeCatalogInvalidParam, "invalid bundle")
		httpresponse.Error(ctx, craneerr)
		return
	}

	catalog.AccountId = account.(auth.Account).ID
	if err := catalogApi.Save(&catalog); err != nil {
		httpresponse.Error(ctx, err)
		return
	}
	httpresponse.Ok(ctx, catalog)
}

func (catalogApi *CatalogApi) DeleteCatalog(ctx *gin.Context) {
	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		log.Error("invalid catalog_id")
		httpresponse.Error(ctx, err)
		return
	}

	cl, err := catalogApi.Get(catalogId)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	if cl.Type == CATALOG_SYSTEM_DEFAULT {
		craneerr := cranerror.NewError(CodeCatalogForbiddenOperation, "forbid update system default")
		httpresponse.Error(ctx, craneerr)
		return
	}

	if err = catalogApi.Delete(catalogId); err != nil {
		log.Errorf("delete catalog error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "delete success")
}

func (catalogApi *CatalogApi) UpdateCatalog(ctx *gin.Context) {
	var catalog Catalog
	if err := ctx.Bind(&catalog); err != nil {
		log.Errorf("invalid param error: %v", err)
		craneerr := cranerror.NewError(CodeCatalogInvalidParam, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if catalog.Bundle == "" {
		craneerr := cranerror.NewError(CodeCatalogInvalidParam, "invalid bundle")
		httpresponse.Error(ctx, craneerr)
		return
	}

	catalogId, err := strconv.ParseUint(ctx.Param("catalog_id"), 10, 64)
	if err != nil {
		log.Error("invalid catalog_id")
		craneerr := cranerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	cl, err := catalogApi.Get(catalogId)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	if cl.Type == CATALOG_SYSTEM_DEFAULT {
		craneerr := cranerror.NewError(CodeCatalogForbiddenOperation, "forbid update system default")
		httpresponse.Error(ctx, craneerr)
		return
	}

	if iconData, err := ImageHandle(ctx.Request); err != nil {
		craneerr := cranerror.NewError(CodeCatalogInvalidIcon, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	} else if iconData == "" {
		catalog.IconData = cl.IconData
	} else {
		catalog.IconData = iconData
	}

	catalog.ID = cl.ID
	catalog.Type = cl.Type
	catalog.AccountId = cl.AccountId
	if err = catalogApi.Update(&catalog); err != nil {
		log.Errorf("update catalog error: %v", err)
		craneerr := cranerror.NewError(CodeCatalogUpdateFaild, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	httpresponse.Ok(ctx, "update success")
}
