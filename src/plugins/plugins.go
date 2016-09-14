package plugins

import (
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/plugins/license"
	rAuthApi "github.com/Dataman-Cloud/crane/src/plugins/registryauth/api"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	log "github.com/Sirupsen/logrus"
)

func Init() {
	conf := config.GetConfig()
	if conf == nil || conf.FeatureFlags == nil {
		log.Warnf("conf or feature flags was nil")
		return
	}

	for _, feature := range conf.FeatureFlags {
		switch feature {
		case apiplugin.License:
			license.Init()
		case apiplugin.RegistryAuth:
			rAuthApi.Init()
		}
	}
}
