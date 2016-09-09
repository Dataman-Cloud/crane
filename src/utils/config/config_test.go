package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFeatureEnabled(t *testing.T) {
	config := &Config{
		FeatureFlags: []string{"foo", "bar"},
	}

	if config.FeatureEnabled("foo") {
		t.Log("feature foo enabled")
	} else {
		t.Error("feature foo should enabled")
	}
}

func TestConfigStruct(t *testing.T) {
	config := new(Config)
	config.CraneAddr = "foobar"
	assert.Equal(t, config.CraneAddr, "foobar")
}

func TestInit(t *testing.T) {
	c := GetConfig()
	assert.NotNil(t, c)
}
