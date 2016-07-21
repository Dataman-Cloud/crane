package account

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/gin-gonic/gin"
)

func (a *AccountApi) GetAccount(ctx *gin.Context) {
	account, err := a.Authenticator.Account(ctx.Param("account_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "1", "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": account})
	}
}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {
	var accountFilter AccountFilter
	accounts, err := a.Authenticator.Accounts(accountFilter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "1", "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": accounts})
	}
}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	token, err := a.Authenticator.Login(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "1", "data": "403"})
		return
	}
	a.TokenStore.Set(token, fmt.Sprintf("%d", acc.ID), time.Now().Add(SESSION_DURATION))
	ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": token})
}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {
	if err := a.TokenStore.Del(ctx.Request.Header.Get("Authorization")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "1", "data": "fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": "success"})
}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {
	var groupFilter GroupFilter
	groups, err := a.Authenticator.Groups(groupFilter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "1", "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": groups})
	}
}

func (a *AccountApi) GetGroup(ctx *gin.Context) {
	groupid, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "1", "data": "bad groupid"})
	}
	group, err := a.Authenticator.Group(groupid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "1", "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": group})
	}
}

func (a *AccountApi) ListGroups(ctx *gin.Context) {
	var groupFilter GroupFilter
	groups, err := a.Authenticator.Groups(groupFilter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "1", "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": "1", "data": groups})
	}
}

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

func (a *AccountApi) GrantServicePermission(ctx *gin.Context) {
	var permission dockerclient.Permission
	if err := ctx.BindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "1", "data": err.Error()})
		return
	}

	err := a.RolexDockerClient.GrantServicePermission(ctx.Param("service_id"), permission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "1", "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": "0", "data": "success"})
}

func (a *AccountApi) RevokeServicePermission(ctx *gin.Context) {
	var permission dockerclient.Permission
	if err := ctx.BindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	err := a.RolexDockerClient.RevokeServicePermission(ctx.Param("service_id"), permission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": "0", "data": "success"})
}
