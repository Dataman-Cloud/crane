package catalog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ExpectedResponse struct {
	Code    int       `json:"code"`
	Data    []Catalog `json:"data"`
	Message string    `json:"message"`
}

func ParseErrorCode(r io.Reader) ExpectedResponse {
	var resp ExpectedResponse
	json.NewDecoder(r).Decode(&resp)
	return resp
}

func TestGetCatalogBadRequest(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalog := MockCatalog()
	catalogApi := NewCatalog(fakeDb.Db)
	catalogApi.DbClient.Save(catalog)

	assert.NotNil(t, catalog.ID)
	assert.Equal(t, catalogApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/:catalog_id", catalogApi.GetCatalog)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", "/test", "malformated"), nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, 15034, ParseErrorCode(w.Body).Code)
}

func TestGetCatalogWrongCatalogID(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalog := MockCatalog()
	catalogApi := NewCatalog(fakeDb.Db)
	catalogApi.DbClient.Save(catalog)

	assert.NotNil(t, catalog.ID)
	assert.Equal(t, catalogApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test/:catalog_id", catalogApi.GetCatalog)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", fmt.Sprintf("%s/%d", "/test", catalog.ID+1), nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, 15031, ParseErrorCode(w.Body).Code)
}

func TestGetCatalogSuccess(t *testing.T) {
	fakeDb := SetupDB()
	catalog := MockCatalog()
	catalogApi := NewCatalog(fakeDb.Db)
	catalogApi.MigriateTable()

	assert.NotNil(t, catalog.ID)
	assert.Equal(t, catalogApi.DbClient, fakeDb.Db)

	//gin.SetMode(gin.TestMode)
	//router := gin.New()
	//router.GET("/test/:catalog_id", catalogApi.GetCatalog)

	//w := httptest.NewRecorder()
	//r, _ := http.NewRequest("GET", fmt.Sprintf("%s/%d", "/test", catalog.ID), nil)
	//router.ServeHTTP(w, r)
	//assert.Equal(t, http.StatusOK, w.Code)
}

func TestListCatalogOK(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalog := MockCatalog()
	catalogApi := NewCatalog(fakeDb.Db)
	catalogApi.DbClient.Save(catalog)

	assert.NotNil(t, catalog.ID)
	assert.Equal(t, catalogApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", catalogApi.ListCatalog)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	// TODO
	//assert.Equal(t, 2, len(ParseErrorCode(w.Body).Data))
}

func TestCreateCatalogOK(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := NewCatalog(fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test", catalogApi.CreateCatalog)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/test", strings.NewReader("FOOBAR"))
	router.ServeHTTP(w, r)
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, 15033, ParseErrorCode(w.Body).Code)
}
