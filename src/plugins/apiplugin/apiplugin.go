package apiplugin

import (
	"github.com/gin-gonic/gin"
)

type CraneApi interface {
	ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc)
}

type ApiPlugin struct {
	Name         string
	Dependencies []string
	Instance     CraneApi
}

var ApiPlugins map[string]*ApiPlugin

func Add(plugin *ApiPlugin) {
	if ApiPlugins == nil {
		ApiPlugins = make(map[string]*ApiPlugin)
	}

	if plugin.Name != "" {
		ApiPlugins[plugin.Name] = plugin
	}

	return
}

const (
	License      = "license"
	Registry     = "registry"
	Account      = "account"
	Catalog      = "catalog"
	Search       = "search"
	RegistryAuth = "registryauth"
	Db           = "db"
)
