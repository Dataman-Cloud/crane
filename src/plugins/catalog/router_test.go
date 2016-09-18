package catalog

import (
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	fakeDb := SetupDB()
	Init(fakeDb.Db)

	assert.NotNil(t, apiplugin.ApiPlugins[apiplugin.Catalog])
}

func TestApiRegister(t *testing.T) {
	router := gin.New()

	fakeDb := SetupDB()
	catalogApi := &CatalogApi{DbClient: fakeDb.Db}
	catalogApi.ApiRegister(router, nil)

	var hasPostCatalogs, hasGetCatalogs, hasGetCatalog, hasPatchCatalog, hasDeleteCatalog bool
	for _, info := range router.Routes() {
		if strings.Contains(info.Path, "/catalogs") && info.Method == "GET" {
			hasGetCatalogs = true
		}

		if strings.Contains(info.Path, "/catalogs") && info.Method == "POST" {
			hasPostCatalogs = true
		}

		if strings.Contains(info.Path, "/catalogs/:catalog_id") && info.Method == "GET" {
			hasGetCatalog = true
		}

		if strings.Contains(info.Path, "/catalogs/:catalog_id") && info.Method == "PATCH" {
			hasPatchCatalog = true
		}

		if strings.Contains(info.Path, "/catalogs/:catalog_id") && info.Method == "DELETE" {
			hasDeleteCatalog = true
		}
	}

	assert.True(t, hasPatchCatalog)
	assert.True(t, hasPostCatalogs)
	assert.True(t, hasGetCatalog)
	assert.True(t, hasGetCatalogs)
	assert.True(t, hasDeleteCatalog)
}
