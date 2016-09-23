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
	"github.com/Dataman-Cloud/crane/src/utils/db"

	log "github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

func Init(conf *config.Config) {
	if conf == nil || conf.FeatureFlags == nil {
		log.Warnf("conf or feature flags was nil")
		return
	}

	for _, feature := range conf.FeatureFlags {
		switch feature {
		case apiplugin.License:
			license.Init(db.DB())
		case apiplugin.RegistryAuth:
			rAuthApi.Init(db.DB())
		case apiplugin.Catalog:
			catalog.Init(db.DB())
		case apiplugin.Registry:
			registry.Init(conf.AccountAuthenticator, conf.RegistryPrivateKeyPath, conf.RegistryAddr)
		case apiplugin.Search:
			search.Init()
		case apiplugin.Account:
			authApi.Init(conf)
		}
	}
}
