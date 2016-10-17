package api

import (
	"os"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"

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
	api := &Api{
		Client: &dockerclient.CraneDockerClient{},
		Config: conf,
	}
	router := api.ApiRouter()

	var hasNetwork bool
	var hasMetrics bool
	var hasStacks bool

	for _, info := range router.Routes() {
		if strings.Contains(info.Path, "nodes") {
			hasMetrics = true
		}
		if strings.Contains(info.Path, "network") {
			hasNetwork = true
		}
		if strings.Contains(info.Path, "stacks") {
			hasStacks = true
		}
	}

	assert.True(t, hasNetwork, "should be true")
	assert.True(t, hasMetrics, "should be true")
	assert.True(t, hasStacks, "should be true")
}
