package plugins

import (
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	authApi "github.com/Dataman-Cloud/crane/src/plugins/auth/api"
	"github.com/Dataman-Cloud/crane/src/plugins/catalog"
	"github.com/Dataman-Cloud/crane/src/plugins/license"
	"github.com/Dataman-Cloud/crane/src/plugins/registry"
	rAuthApi "github.com/Dataman-Cloud/crane/src/plugins/registryauth/api"
	"github.com/Dataman-Cloud/crane/src/plugins/search"
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
		case apiplugin.Catalog:
			catalog.Init()
		case apiplugin.Registry:
			registry.Init()
		case apiplugin.Search:
			search.Init()
		case apiplugin.Account:
			authApi.Init()
		}
	}
}
