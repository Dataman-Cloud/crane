package api

import (
	"testing"
	"strings"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/stretchr/testify/assert"
)

func TestApiRouter(t *testing.T) {
	conf := config.GetConfig()
	conf.FeatureFlags = []string{"logging", "account", "search"}
	api := &Api{
		Client: &dockerclient.CraneDockerClient{},
		Config: conf,
	}
	router := api.ApiRouter()

	var hasAuth bool
	var hasNetwork bool
	var hasMetrics bool
	var hasStacks bool
	var hasSearch bool

	for _, info := range router.Routes() {
		if strings.Contains(info.Path, "nodes") {
			hasMetrics = true
		}
		if strings.Contains(info.Path, "network") {
			hasNetwork = true
		}
		if strings.Contains(info.Path, "accounts") {
			hasAuth = true
		}
		if strings.Contains(info.Path, "stacks") {
			hasStacks = true
		}
		if strings.Contains(info.Path, "search") {
			hasSearch = true
		}
	}

	assert.True(t, hasAuth, "should be true")
	assert.True(t, hasNetwork, "should be true")
	assert.True(t, hasMetrics, "should be true")
	assert.True(t, hasStacks, "should be true")
	assert.True(t, hasSearch, "should be true")
}
