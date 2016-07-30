package config

import (
	"testing"

	"github.com/Dataman-Cloud/rolex/src/util/config"
)

func TestConfigFeatureEnabled(t *testing.T) {
	config := &config.Config{
		FeatureFlags: []string{"foo", "bar"},
	}

	if config.FeatureEnabled("foo") {
		t.Log("feature foo enabled")
	} else {
		t.Error("feature foo should enabled")
	}
}
