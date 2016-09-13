package api

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/plugins/auth/authenticators"
	chains "github.com/Dataman-Cloud/crane/src/plugins/auth/middlewares"
	"github.com/Dataman-Cloud/crane/src/plugins/auth/token_store"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config            *config.Config
	Authenticator     auth.Authenticator
	TokenStore        auth.TokenStore
	CraneDockerClient *dockerclient.CraneDockerClient
	Authorization     gin.HandlerFunc
}

func (account *AccountApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	if account.Config.AccountTokenStore == "default" {
		account.TokenStore = token_store.NewDefaultStore()
	} else if account.Config.AccountTokenStore == "cookie_store" {
		account.TokenStore = token_store.NewCookieStore()
	}

	if account.Config.AccountAuthenticator == "default" {
		account.Authenticator = authenticators.NewDefaultAuthenticator()
	} else if account.Config.AccountAuthenticator == "db" {
		account.Authenticator = authenticators.NewDBAuthenticator()
	}

	account.Authorization = chains.Authorization(account.TokenStore, account.Authenticator)
	AuthorizeServiceAccess := chains.AuthorizeServiceAccess()

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

	serviceV1 := router.Group("/api/v1/")
	{
		serviceV1.Use(middlewares...)
		serviceV1.POST("services/:service_id/permissions", AuthorizeServiceAccess(auth.PermAdmin),
			account.GrantServicePermission)
		serviceV1.DELETE("services/:service_id/permissions/:permission_id", AuthorizeServiceAccess(auth.PermAdmin),
			account.RevokeServicePermission)
	}
}
