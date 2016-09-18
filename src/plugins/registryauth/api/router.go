package api

import (
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	rauth "github.com/Dataman-Cloud/crane/src/plugins/registryauth"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type RegistryAuthApi struct{}

func Init(dbClient *gorm.DB) {
	log.Infof("begin to init and enable plugin: %s", apiplugin.RegistryAuth)
	rauth.Init(dbClient)

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.RegistryAuth,
		Dependencies: []string{apiplugin.Db, apiplugin.Account},
		Instance:     &RegistryAuthApi{},
	}

	apiplugin.Add(apiPlugin)
	log.Infof("init and enable plugin: %s success", apiplugin.RegistryAuth)
}

func (api *RegistryAuthApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	rauthv1 := router.Group("/registryauth/v1")
	if middlewares != nil {
		for _, middleware := range middlewares {
			rauthv1.Use(middleware)
		}
	}
	{
		rauthv1.POST("/registryauths", api.Create)
		rauthv1.GET("/registryauths", api.List)
		rauthv1.DELETE("/registryauths/:rauth_name", api.Delete)
	}
}
