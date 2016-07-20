package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AccountApi) GetAccount(ctx *gin.Context) {}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {}

func (a *AccountApi) CreateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "1", "data": "403"})
		return
	}
}

func (a *AccountApi) UpdateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "1", "data": "403"})
		return
	}
}

func (a *AccountApi) DeleteGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "1", "data": "403"})
		return
	}
}

func (a *AccountApi) JoinGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "1", "data": "403"})
		return
	}
}

func (a *AccountApi) LeaveGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "1", "data": "403"})
		return
	}
}
