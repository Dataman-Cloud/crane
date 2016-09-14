package registry

import (
	"os"

	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Init() {
	log.Infof("begin to init and enable plugin: %s", apiplugin.Registry)
	conf := config.GetConfig()
	if conf == nil {
		log.Errorf("init plugin: %s failed conf is nil", apiplugin.Registry)
		os.Exit(1)
	}

	registryApi := NewRegistry(conf.AccountAuthenticator, conf.RegistryPrivateKeyPath, conf.RegistryAddr)

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.Registry,
		Dependencies: []string{apiplugin.Db},
		Instance:     registryApi,
	}

	apiplugin.Add(apiPlugin)
	log.Infof("init and enable plugin: %s success", apiplugin.License)
}

func (registry *Registry) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	registryV1 := router.Group("/registry/v1")
	{
		registryV1.GET("/token", registry.Token)
		registryV1.POST("/notifications", registry.Notifications)
	}

	registryV1Protected := router.Group("/registry/v1", middlewares...)
	{
		registryV1Protected.GET("/repositories/mine", registry.MineRepositories)
		registryV1Protected.GET("/repositories/public", registry.PublicRepositories) // under library or tag marked as public
		registryV1Protected.GET("/tag/list/:namespace/:image", registry.TagList)
		registryV1Protected.GET("/manifests/:reference/:namespace/:image", registry.GetManifests)
		registryV1Protected.PATCH("/:namespace/:image/publicity", registry.ImagePublicity)
		registryV1Protected.DELETE("/manifests/:reference/:namespace/:image", registry.DeleteManifests)
	}
}
