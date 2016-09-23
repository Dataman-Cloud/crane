package api

import (
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestApiRouter(t *testing.T) {
	conf := config.GetConfig()
	conf.FeatureFlags = []string{"logging", "account", "search"}
	api := &AccountApi{
		Config: conf,
	}

	router := gin.New()

	api.ApiRegister(router)

	var hasNetwork bool
	var hasMetrics bool
	var hasStacks bool

	for _, info := range router.Routes() {
		if strings.Contains(info.Path, "accounts") {
			hasMetrics = true
		}
		if strings.Contains(info.Path, "aboutme") {
			hasNetwork = true
		}
		if strings.Contains(info.Path, "groups") {
			hasStacks = true
		}
	}

	assert.True(t, hasNetwork, "should be true")
	assert.True(t, hasMetrics, "should be true")
	assert.True(t, hasStacks, "should be true")
}
