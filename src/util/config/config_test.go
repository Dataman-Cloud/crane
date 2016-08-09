package config

import (
	"testing"
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
