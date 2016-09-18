package api

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/plugins/auth/authenticators"
	chains "github.com/Dataman-Cloud/crane/src/plugins/auth/middlewares"
	"github.com/Dataman-Cloud/crane/src/plugins/auth/token_store"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config            *config.Config
	Authenticator     auth.Authenticator
	TokenStore        auth.TokenStore
	CraneDockerClient *dockerclient.CraneDockerClient
	Authorization     gin.HandlerFunc
}

func Init(conf *config.Config) {
	log.Infof("begin to init and enable plugin: %s", apiplugin.Account)
	accountApi := &AccountApi{Config: conf}
	if conf.AccountTokenStore == "default" {
		accountApi.TokenStore = token_store.NewDefaultStore()
	} else if conf.AccountTokenStore == "cookie_store" {
		accountApi.TokenStore = token_store.NewCookieStore()
	}

	if conf.AccountAuthenticator == "default" {
		accountApi.Authenticator = authenticators.NewDefaultAuthenticator()
	} else if conf.AccountAuthenticator == "db" {
		accountApi.Authenticator = authenticators.NewDBAuthenticator()
	}

	accountApi.Authorization = chains.Authorization(accountApi.TokenStore, accountApi.Authenticator)

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.Account,
		Dependencies: []string{apiplugin.Db},
		Instance:     accountApi,
	}

	apiplugin.Add(apiPlugin)
	log.Infof("init and enable plugin: %s success", apiplugin.Account)
}

func (account *AccountApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	accountV1 := router.Group("/account/v1")
	{
		accountV1.Use(account.Authorization)
		accountV1.Use(middlewares...)
		accountV1.GET("/aboutme", account.GetAccountInfo)
		accountV1.GET("/accounts/:account_id", account.GetAccount)
		accountV1.GET("/accounts", account.ListAccounts)

		accountV1.POST("/logout", account.AccountLogout)

		accountV1.GET("/accounts/:account_id/groups", account.AccountGroups)
		accountV1.POST("/accounts/:account_id/groups/:group_id", account.JoinGroup)
		accountV1.DELETE("/accounts/:account_id/groups/:group_id", account.LeaveGroup)

		accountV1.GET("/groups", account.ListGroups)
		accountV1.POST("/groups", account.CreateGroup)
		accountV1.PATCH("/groups", account.UpdateGroup)
		accountV1.GET("/groups/:group_id", account.GetGroup)
		accountV1.DELETE("/groups/:group_id", account.DeleteGroup)
		accountV1.POST("/groups/:group_id/account", account.CreateAccount)
		accountV1.GET("/groups/:group_id/accounts", account.GetGroup)
	}

	router.POST("/account/v1/login", account.AccountLogin)

	AuthorizeServiceAccess := chains.AuthorizeServiceAccess()
	serviceV1 := router.Group("/api/v1/")
	{
		serviceV1.Use(middlewares...)
		serviceV1.POST("services/:service_id/permissions", AuthorizeServiceAccess(auth.PermAdmin),
			account.GrantServicePermission)
		serviceV1.DELETE("services/:service_id/permissions/:permission_id", AuthorizeServiceAccess(auth.PermAdmin),
			account.RevokeServicePermission)
	}
}
