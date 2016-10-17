package registry

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/db"

	_ "github.com/erikstmartin/go-testdb"
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

func FakeAuthenticate(ctx *gin.Context) {
	ctx.Set("account", auth.Account{ID: 1})
	ctx.Next()
}

func TestMain(m *testing.M) {
	server = startHttpServer()
	baseUrl = server.URL
	defer server.Close()
	os.Exit(m.Run())
}

func startHttpServer() *httptest.Server {
	router := gin.New()

	registryV1 := router.Group("/registry/v1")
	{
		registryV1.GET("/token", r.Token)
		registryV1.POST("/notifications", r.Notifications)
	}

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
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest, "response status code should be equal")

	r.PrivateKeyPath = ""
	req, err = http.NewRequest("GET", baseUrl+"/registry/v1/token", nil)
	req.SetBasicAuth("admin@admin.com", "adminadmin")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
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
