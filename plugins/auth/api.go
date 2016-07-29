package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Dataman-Cloud/rolex/model"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (a *AccountApi) CreateAccount(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		log.Errorf("create account error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	if acc.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "password can not be null"})
		return
	}

	if acc.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "email can not be null"})
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "invalid groupid"})
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	acc.LoginAt = time.Now()
	if err := a.Authenticator.CreateAccount(groupId, &acc); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "create success"})
}

func (a *AccountApi) GetAccountInfo(ctx *gin.Context) {
	account, _ := ctx.Get("account")
	account, err := a.Authenticator.Account(account.(Account).ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": account})
	}
}

func (a *AccountApi) GetAccount(ctx *gin.Context) {
	account, err := a.Authenticator.Account(ctx.Param("account_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": account})
	}
}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")

	accounts, err := a.Authenticator.Accounts(listOptions.(model.ListOptions))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": "404"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": accounts})
	}
}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		log.Errorf("login error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	token, err := a.Authenticator.Login(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "403"})
		return
	}
	a.TokenStore.Set(ctx, token, fmt.Sprintf("%d", acc.ID), time.Now().Add(SESSION_DURATION))
	acc.Password = ""
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": token})
}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {
	if err := a.TokenStore.Del(ctx, ctx.Request.Header.Get("Authorization")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": "fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (a *AccountApi) GroupAccounts(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64); err != nil {
		log.Errorf("invalid groupid: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"group_id": groupId,
		}
	}

	accounts, err := a.Authenticator.GroupAccounts(listOptions)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": accounts})
	}
}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64); err != nil {
		log.Errorf("invalid accountid: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"account_id": accountId,
		}
	}

	groups, err := a.Authenticator.AccountGroups(listOptions)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": groups})
	}
}

func (a *AccountApi) GetGroup(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "bad groupid"})
		return
	}
	group, err := a.Authenticator.Group(groupId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": group})
	}
}

func (a *AccountApi) ListGroups(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")
	groups, err := a.Authenticator.Groups(listOptions.(model.ListOptions))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": groups})
	}
}

func (a *AccountApi) CreateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "403"})
		return
	}

	var group Group
	if err := ctx.BindJSON(&group); err != nil {
		log.Errorf("create group request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	account, _ := ctx.Get("account")
	group.CreaterId = account.(Account).ID
	if err := a.Authenticator.CreateGroup(&group); err != nil {
		log.Errorf("create group db operation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": "create success"})
}

func (a *AccountApi) UpdateGroup(ctx *gin.Context) {
	var group Group

	if err := ctx.BindJSON(&group); err != nil {
		log.Errorf("update group request body parse json error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "403"})
		return
	}

	if err := a.Authenticator.UpdateGroup(&group); err != nil {
		log.Errorf("update group db operation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": "update success"})
}

func (a *AccountApi) DeleteGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "403"})
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		log.Errorf("delete group invalid groupId error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "invalid group_id"})
	}

	if err := a.Authenticator.DeleteGroup(groupId); err != nil {
		log.Errorf("delete group db operation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": "delete success"})
}

func (a *AccountApi) JoinGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "403"})
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "bad accountid"})
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "bad accountid"})
		return
	}

	if err := a.Authenticator.JoinGroup(accountId, groupId); err != nil {
		log.Errorf("user join group db operation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": "user join group success"})
}

func (a *AccountApi) LeaveGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "403"})
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "bad accountid"})
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "bad groupid"})
		return
	}

	if err := a.Authenticator.LeaveGroup(accountId, groupId); err != nil {
		log.Errorf("user leave  group db operation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"code": 0, "data": "user leave group success"})
}

func (a *AccountApi) GrantServicePermission(ctx *gin.Context) {
	var param struct {
		GroupID uint64 `json:"GroupID"`
		Perm    string `json:"Perm"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": err.Error()})
		return
	}

	err := a.RolexDockerClient.ServiceAddLabel(ctx.Param("service_id"), PermissionGrantLabelsPairFromGroupIdAndPerm(param.GroupID, param.Perm))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (a *AccountApi) RevokeServicePermission(ctx *gin.Context) {
	permission_id := ctx.Param("permission_id")

	if len(strings.SplitN(permission_id, "-", 2)) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "permission id not valid"})
		return
	}

	labels := PermissionRevokeLabelKeysFromPermissionId(permission_id)

	err := a.RolexDockerClient.ServiceRemoveLabel(ctx.Param("service_id"), labels)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}
