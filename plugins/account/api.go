package account

import (
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
	Config config.Config
}

func (a *AccountApi) GetAccount(ctx *gin.Context)    {}
func (a *AccountApi) ListAccounts(ctx *gin.Context)  {}
func (a *AccountApi) AccountLogin(ctx *gin.Context)  {}
func (a *AccountApi) AccountLogout(ctx *gin.Context) {}
func (a *AccountApi) AccountAcls(ctx *gin.Context)   {}
func (a *AccountApi) AccountRoles(ctx *gin.Context)  {}
func (a *AccountApi) CreateRole(ctx *gin.Context)    {}
func (a *AccountApi) UpdateRole(ctx *gin.Context)    {}
func (a *AccountApi) DeleteRole(ctx *gin.Context)    {}
func (a *AccountApi) RoleAcls(ctx *gin.Context)      {}
