package account

import (
	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config config.Config
}

func (a *AccountApi) RegisterApiForAccount(router *gin.Engine) {
	accountV1 := router.Group("/account/v1")
	{
		accountV1.GET("accounts/:account_id", account.GetAccount)
		accountV1.GET("accounts", account.ListAccounts)
		accountV1.GET("accounts/:account_id/roles", account.AccountRoles)
		accountV1.GET("accounts/:account_id/acls", account.AccountAcls)
		accountV1.POST("accounts/:account_id/login", account.AccountLogin)
		accountV1.POST("accounts/:account_id/logout", account.AccountLogout)
		accountV1.POST("roles", account.CreateRole)
		accountV1.PATCH("roles/:role_id", account.UpdateRole)
		accountV1.DELETE("roles/:role_id", account.DeleteRole)
		accountV1.GET("roles/:role_id/acls", account.RoleAcls)
	}
}
