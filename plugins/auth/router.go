package auth

import (
	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config            *config.Config
	Authenticator     Authenticator
	TokenStore        TokenStore
	RolexDockerClient *dockerclient.RolexDockerClient
}

func (account *AccountApi) RegisterApiForAccount(router *gin.Engine, authorization ...gin.HandlerFunc) {
	accountV1 := router.Group("/account/v1")
	{
		accountV1.Use(authorization...)
		accountV1.GET("/accounts/:account_id", account.GetAccount)
		accountV1.GET("/accounts", account.ListAccounts)

		accountV1.POST("/logout", account.AccountLogout)

		accountV1.GET("/accounts/:account_id/groups", account.AccountGroups)
		accountV1.POST("/accounts/:account_id/groups/:group_id", account.JoinGroup)
		accountV1.DELETE("/accounts/:account_id/groups/:group_id", account.LeaveGroup)

		accountV1.GET("/groups", account.ListGroups)
		accountV1.GET("/groups/:group_id", account.GetGroup)
		accountV1.POST("/groups", account.CreateGroup)
		accountV1.PATCH("/groups", account.UpdateGroup)
		accountV1.DELETE("/groups/:group_id", account.DeleteGroup)
	}

	router.POST("/account/v1/login", account.AccountLogin)
	router.POST("/account/v1/accounts", account.CreateAccount)

	serviceV1 := router.Group("/api/v1/")
	{
		serviceV1.Use(authorization...)
		serviceV1.POST("services/:service_id/permissions", account.GrantServicePermission)
		serviceV1.DELETE("services/:service_id/permissions/:permission", account.RevokeServicePermission)
	}
}
