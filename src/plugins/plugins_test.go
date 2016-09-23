package plugins

import (
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"
)

func TestInitNilConfig(t *testing.T) {
	Init(nil)
}

func TestInitSearchConfig(t *testing.T) {
	conf := config.GetConfig()
	conf.FeatureFlags = []string{"search"}
	Init(conf)
}
