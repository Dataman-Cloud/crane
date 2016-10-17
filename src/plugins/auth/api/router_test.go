package api

import (
	"os"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRouter(t *testing.T) {
	os.Setenv("CRANE_ADDR", "foobar")
	os.Setenv("CRANE_SWARM_MANAGER_IP", "foobar")
	os.Setenv("CRANE_DOCKER_CERT_PATH", "foobar")
	os.Setenv("CRANE_DB_DRIVER", "foobar")
	os.Setenv("CRANE_DB_DSN", "foobar")
	os.Setenv("CRANE_FEATURE_FLAGS", "foobar")
	os.Setenv("CRANE_REGISTRY_PRIVATE_KEY_PATH", "foobar")
	os.Setenv("CRANE_REGISTRY_ADDR", "foobar")
	os.Setenv("CRANE_ACCOUNT_AUTHENTICATOR", "foobar")
	defer os.Setenv("CRANE_ADDR", "")
	defer os.Setenv("CRANE_SWARM_MANAGER_IP", "")
	defer os.Setenv("CRANE_DOCKER_CERT_PATH", "")
	defer os.Setenv("CRANE_DB_DRIVER", "")
	defer os.Setenv("CRANE_DB_DSN", "")
	defer os.Setenv("CRANE_FEATURE_FLAGS", "")
	defer os.Setenv("CRANE_REGISTRY_PRIVATE_KEY_PATH", "")
	defer os.Setenv("CRANE_REGISTRY_ADDR", "")
	defer os.Setenv("CRANE_ACCOUNT_AUTHENTICATOR", "")

	conf := config.InitConfig()
	conf.FeatureFlags = []string{"logging", "account", "search"}
	api := &AccountApi{
		Config: conf,
	}

	router := gin.New()

	api.ApiRegister(router)

	// TODO (weitao) improve me by func assert.Contains(t, ) . refer plugins/registry/router_test.go
	var hasAccounts bool
	var hasAboutme bool
	var hasGroups bool

	for _, info := range router.Routes() {
		if strings.Contains(info.Path, "accounts") {
			hasAccounts = true
		}
		if strings.Contains(info.Path, "aboutme") {
			hasAboutme = true
		}
		if strings.Contains(info.Path, "groups") {
			hasGroups = true
		}
	}

	assert.True(t, hasAboutme, "should be true")
	assert.True(t, hasAccounts, "should be true")
	assert.True(t, hasGroups, "should be true")
}
