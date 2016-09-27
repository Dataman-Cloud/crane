package registry

import (
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/db"

	"github.com/erikstmartin/go-testdb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRegistry(t *testing.T) {
	dbDriver := "testdb"
	dbDSN := ""
	dbClient, err := db.NewDB(dbDriver, dbDSN)
	assert.Nil(t, err)

	// TODO (wtzhou) how to assert dbClient
	registry := NewRegistry("db", "", "", dbDriver, dbDSN, dbClient)
	assert.Equal(t, "db", registry.AccountAuthenticator)
	assert.Equal(t, "", registry.PrivateKeyPath)
}

func TestMigrateTable(t *testing.T) {
	dbDriver := "testdb"
	dbDSN := ""
	dbClient, err := db.NewDB(dbDriver, dbDSN)
	assert.Nil(t, err)

	// TODO (wtzhou) how to assert dbClient
	registry := NewRegistry("db", "", "", dbDriver, dbDSN, dbClient)
	assert.Equal(t, "db", registry.AccountAuthenticator)
	assert.Equal(t, "", registry.PrivateKeyPath)
	registry.migrateTable()
	assert.Nil(t, nil)
}

// TODO (wtzhou) refactor me by test/testing mock
var (
	baseUrl string
	server  *httptest.Server
	r       = &Registry{
		AccountAuthenticator: "auth_type",
		PrivateKeyPath:       "private_key.test",
		RegistryAddr:         "registry_addr",
		Authenticator:        auth.NewMockAuthenticator(),
	}
)

func TestMain(m *testing.M) {
	server = startHttpServer()
	baseUrl = server.URL
	r.RegistryAddr = baseUrl
	defer server.Close()
	os.Exit(m.Run())
}

func FakeAuthenticate(ctx *gin.Context) {
	if _, _, ok := ctx.Request.BasicAuth(); ok {
		ctx.Set("account", auth.ReferenceToValue(&auth.Account{
			ID:       1,
			Title:    "",
			Email:    "admin@admin.com",
			Phone:    "",
			Password: "adminadmin",
		}))
	}
	ctx.Next()
}

func startHttpServer() *httptest.Server {
	dbDriver := "testdb"
	dbDSN := ""
	dbClient, _ := db.NewDB(dbDriver, dbDSN)
	r.DbClient = dbClient

	router := gin.New()
	registryV1 := router.Group("/registry/v1", FakeAuthenticate)
	{
		registryV1.GET("/token", r.Token)
		registryV1.POST("/notifications", r.Notifications)
		registryV1.DELETE("/manifests/:namespace/:image", r.DeleteManifests)
		//registry_addr/v2/admin/2048/manifests/Tim
	}
	router.DELETE("/v2/:namespace/:image/manifests/:digest", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"test": "success",
		})
	})

	return httptest.NewServer(router)
}

func TestToken(t *testing.T) {
	req, err := http.NewRequest("GET", baseUrl+"/registry/v1/token", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response status code should be equal")

	// private key path lost
	// assert make token failed
	req, err = http.NewRequest("GET", baseUrl+"/registry/v1/token", nil)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Log("pass")
	}
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest, "response status code should be equal")

	r.PrivateKeyPath = ""
	req, err = http.NewRequest("GET", baseUrl+"/registry/v1/token", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Log("pass")
	}
	assert.Equal(t, resp.StatusCode, http.StatusServiceUnavailable, "response status code should be equal")
}

func TestNotifications(t *testing.T) {
	notification := Notification{
		Events: []Event{
			{
				ID: "test",
				Target: &Target{
					MediaType: "application",
				},
			},
		},
	}
	body, err := json.Marshal(notification)
	if err != nil {
		t.Log(err)
		return
	}

	req, err := http.NewRequest("POST", baseUrl+"/registry/v1/notifications", strings.NewReader(string(body)))
	req.Header.Add("User-Agent", "Go-http-client")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response status code should be equal")
}

func TestDeleteManifests(t *testing.T) {
	req, err := http.NewRequest("DELETE", baseUrl+"/registry/v1/manifests/admin/2048", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusUnauthorized, "response status code should be equal")

	req, err = http.NewRequest("DELETE", baseUrl+"/registry/v1/manifests/admin/2048", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusServiceUnavailable, "response status code should be equal")

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "digest", "tag", "namespace", "image"}
		result := `1,Tim,latest,admin,2048
			   3,Bob,30,admin,2048`
		return testdb.RowsFromCSVString(columns, result), nil
	})
	dbDriver := "testdb"
	dbDSN := ""
	dbClient, _ := db.NewDB(dbDriver, dbDSN)
	r.DbClient = dbClient
	req, err = http.NewRequest("DELETE", baseUrl+"/registry/v1/manifests/admin/2048?tag=test", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response status code should be equal")

	req, err = http.NewRequest("DELETE", baseUrl+"/registry/v1/manifests/admin/2048?namespace=admin&image=2048", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response status code should be equal")
}
