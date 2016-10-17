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

func Init(conf *config.Config) error {
	if conf == nil || conf.FeatureFlags == nil {
		log.Warnf("conf or feature flags was nil")
		return nil
	}

	for _, feature := range conf.FeatureFlags {
		switch feature {
		case apiplugin.License:
			// TODO (wtzhou) improve me: how to merge the following db.NewDB into single
			dbClient, err := db.NewDB(conf.DbDriver, conf.DbDSN)
			if err != nil {
				return err
			}
			license.Init(dbClient)
		case apiplugin.RegistryAuth:
			dbClient, err := db.NewDB(conf.DbDriver, conf.DbDSN)
			if err != nil {
				return err
			}
			rAuthApi.Init(dbClient)
		case apiplugin.Catalog:
			dbClient, err := db.NewDB(conf.DbDriver, conf.DbDSN)
			if err != nil {
				return err
			}
			catalog.Init(dbClient)
		case apiplugin.Registry:
			dbClient, err := db.NewDB(conf.DbDriver, conf.DbDSN)
			if err != nil {
				return err
			}
			registry.Init(conf.AccountAuthenticator, conf.RegistryPrivateKeyPath, conf.RegistryAddr, conf.DbDriver, conf.DbDSN, dbClient)
		case apiplugin.Search:
			search.Init()
		case apiplugin.Account:
			authApi.Init(conf)
		}
	}
	return nil
}
