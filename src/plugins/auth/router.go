package auth

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config            *config.Config
	Authenticator     Authenticator
	TokenStore        TokenStore
	RolexDockerClient *dockerclient.RolexDockerClient
}

func (account *AccountApi) RegisterApiForAccount(router *gin.Engine,
	authorizeMiddlwares map[string](func(permissionRequired Permission) gin.HandlerFunc),
	middlewares ...gin.HandlerFunc) {
	accountV1 := router.Group("/account/v1")
	{
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
		accountV1.POST("/account/v1/accounts", account.CreateAccount)
		accountV1.GET("/groups/:group_id/accounts", account.GetGroup)
	}

	router.POST("/account/v1/login", account.AccountLogin)

	serviceV1 := router.Group("/api/v1/")
	{
		serviceV1.Use(middlewares...)
		serviceV1.POST("services/:service_id/permissions",
			authorizeMiddlwares["AuthorizeServiceAccess"](PermAdmin),
			account.GrantServicePermission)
		serviceV1.DELETE("services/:service_id/permissions/:permission_id",
			authorizeMiddlwares["AuthorizeServiceAccess"](PermAdmin),
			account.RevokeServicePermission)
	}
}
