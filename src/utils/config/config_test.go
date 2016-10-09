package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFeatureEnabled(t *testing.T) {
	config := &Config{
		FeatureFlags: []string{"foo", "bar"},
	}

	assert.True(t, config.FeatureEnabled("foo"), "feature foo should be enabled")
}

func TestConfigStruct(t *testing.T) {
	config := new(Config)
	config.CraneAddr = "foobar"
	assert.Equal(t, config.CraneAddr, "foobar")
}

func TestInitConfig(t *testing.T) {
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

	config = new(Config)
	c := InitConfig()
	assert.NotNil(t, c)
	t.Logf("config struct: %+v", c)
}
