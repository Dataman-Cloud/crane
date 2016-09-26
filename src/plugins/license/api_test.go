package license

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"strings"
)

type FakeDB struct {
	FilePath string
	Db       *gorm.DB
}

func SetupDB() *FakeDB {
	sqliteDbPath := filepath.Join(os.TempDir(), "test.db")
	db, _ := gorm.Open("sqlite3", sqliteDbPath)
	db.AutoMigrate(&Setting{})
	return &FakeDB{FilePath: sqliteDbPath, Db: db}
}

func TearDownDB(fakeDb *FakeDB) {
	fakeDb.Db.Close()
	os.Remove(fakeDb.FilePath)
}

type ExpectedResponse struct {
	Code    int       `json:"code"`
	Data    []Setting `json:"data"`
	Message string    `json:"message"`
}

func ParseErrorCode(r io.Reader) ExpectedResponse {
	var resp ExpectedResponse
	json.NewDecoder(r).Decode(&resp)
	return resp
}

func MockLicense() *Setting {
	return &Setting{
		ID:      1,
		License: "dsfsdfsd",
	}
}

func TestCreate(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	licenseApi := &LicenseApi{DbClient: fakeDb.Db}
	license := MockLicense()
	licenseApi.DbClient.Save(license)

	assert.NotNil(t, license.ID)
	assert.Equal(t, licenseApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/test", licenseApi.Create)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/test", strings.NewReader(`{ID:555, License: "3sdfrsdfsd"}`))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, 16002, ParseErrorCode(w.Body).Code)
}

func TestGet(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	licenseApi := &LicenseApi{DbClient: fakeDb.Db}
	license := MockLicense()
	licenseApi.DbClient.Save(license)

	assert.NotNil(t, license.ID)
	assert.Equal(t, licenseApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/test", licenseApi.Get)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.Equal(t, 16003, ParseErrorCode(w.Body).Code)
}

func TestGetLicenseValidity(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	licenseApi := &LicenseApi{DbClient: fakeDb.Db}
	license := MockLicense()
	licenseApi.DbClient.Save(license)

	assert.NotNil(t, license.ID)
	assert.Equal(t, licenseApi.DbClient, fakeDb.Db)

	gin.SetMode(gin.TestMode)

	_, err := licenseApi.GetLicenseValidity()
	assert.NotNil(t, err)
}
