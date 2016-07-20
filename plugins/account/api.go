package account

import (
	"github.com/gin-gonic/gin"
)

func (a *Account) GetAccount(ctx *gin.Context)    {}
func (a *Account) ListAccounts(ctx *gin.Context)  {}
func (a *Account) AccountLogin(ctx *gin.Context)  {}
func (a *Account) AccountLogout(ctx *gin.Context) {}
func (a *Account) AccountAcls(ctx *gin.Context)   {}
func (a *Account) AccountRoles(ctx *gin.Context)  {}
func (a *Account) CreateRole(ctx *gin.Context)    {}
func (a *Account) UpdateRole(ctx *gin.Context)    {}
func (a *Account) DeleteRole(ctx *gin.Context)    {}
func (a *Account) RoleAcls(ctx *gin.Context)      {}
