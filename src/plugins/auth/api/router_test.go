package api

import (
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRouter(t *testing.T) {
	conf := config.GetConfig()
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
