package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"
	"github.com/Dataman-Cloud/crane/src/utils/model"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (a *AccountApi) CreateAccount(ctx *gin.Context) {
	var acc auth.Account
	if err := ctx.BindJSON(&acc); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if acc.Password == "" {
		craneerr := cranerror.NewError(auth.CodeAccountCreateParamError, "password can not be null")
		httpresponse.Error(ctx, craneerr)
		return
	}

	if acc.Email == "" {
		craneerr := cranerror.NewError(auth.CodeAccountCreateParamError, "email can not be null")
		httpresponse.Error(ctx, craneerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateParamError, "invalid groupid")
		httpresponse.Error(ctx, craneerr)
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	acc.LoginAt = time.Now()
	if err := a.Authenticator.CreateAccount(groupId, &acc); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateAuthenticatorError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) GetAccountInfo(ctx *gin.Context) {
	account, _ := ctx.Get("account")
	account, err := a.Authenticator.Account(account.(auth.Account).ID)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGetAccountNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, account)
	}
}

func (a *AccountApi) GetAccount(ctx *gin.Context) {
	account, err := a.Authenticator.Account(ctx.Param("account_id"))
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGetAccountNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, account)
	}
}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")

	accounts, err := a.Authenticator.Accounts(listOptions.(model.ListOptions))
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGetAccountNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, accounts)
	}
}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {
	var acc auth.Account
	if err := ctx.BindJSON(&acc); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountLoginParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	token, err := a.Authenticator.Login(&acc)
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	a.TokenStore.Set(ctx, token, fmt.Sprintf("%d", acc.ID), time.Now().Add(auth.SESSION_DURATION))
	httpresponse.Ok(ctx, token)
}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {
	if err := a.TokenStore.Del(ctx, ctx.Request.Header.Get("Authorization")); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountLogoutError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) GroupAccounts(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64); err != nil {
		log.Errorf("invalid groupid: %v", err)
		craneerr := cranerror.NewError(auth.CodeAccountGroupAccountsGroupIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"group_id": groupId,
		}
	}

	accounts, err := a.Authenticator.GroupAccounts(listOptions)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGroupAccountsNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, accounts)
	}
}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountAccoutGroupsAccountIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"account_id": accountId,
		}
	}

	groups, err := a.Authenticator.AccountGroups(listOptions)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountAccoutGroupsNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, groups)
	}
}

func (a *AccountApi) GetGroup(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGetGroupGroupIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	group, err := a.Authenticator.Group(groupId)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGetGroupGroupIdNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, group)
	}
}

func (a *AccountApi) ListGroups(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")
	groups, err := a.Authenticator.Groups(listOptions.(model.ListOptions))
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountListGroupNotFoundError, err.Error())
		httpresponse.Error(ctx, craneerr)
	} else {
		httpresponse.Ok(ctx, groups)
	}
}

func (a *AccountApi) CreateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		craneerr := cranerror.NewError(auth.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		httpresponse.Error(ctx, craneerr)
		return
	}

	var group auth.Group
	if err := ctx.BindJSON(&group); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	account, _ := ctx.Get("account")
	group.CreaterId = account.(auth.Account).ID
	if err := a.Authenticator.CreateGroup(&group); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateGroupFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) UpdateGroup(ctx *gin.Context) {
	var group auth.Group

	if !a.Authenticator.ModificationAllowed() {
		craneerr := cranerror.NewError(auth.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		httpresponse.Error(ctx, craneerr)
		return
	}

	if err := ctx.BindJSON(&group); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountCreateGroupParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if err := a.Authenticator.UpdateGroup(&group); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountUpdateGroupParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) DeleteGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		craneerr := cranerror.NewError(auth.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		httpresponse.Error(ctx, craneerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountDeleteGroupGroupIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
	}

	if err := a.Authenticator.DeleteGroup(groupId); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountDeleteGroupFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) JoinGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		craneerr := cranerror.NewError(auth.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		httpresponse.Error(ctx, craneerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountJoinGroupGroupIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountJoinGroupAccountIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if err := a.Authenticator.JoinGroup(accountId, groupId); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountJoinGroupFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) LeaveGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		craneerr := cranerror.NewError(auth.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		httpresponse.Error(ctx, craneerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountLeaveGroupAccountIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountLeaveGroupGroupIdNotValidError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	if err := a.Authenticator.LeaveGroup(accountId, groupId); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountLeaveGroupFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return

	}

	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) GrantServicePermission(ctx *gin.Context) {
	var param struct {
		GroupID uint64 `json:"GroupID"`
		Perm    string `json:"Perm"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGrantServicePermissionParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	err := a.CraneDockerClient.ServiceAddLabel(ctx.Param("service_id"), auth.PermissionGrantLabelsPairFromGroupIdAndPerm(param.GroupID, param.Perm))
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountGrantServicePermissionFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}

func (a *AccountApi) RevokeServicePermission(ctx *gin.Context) {
	permissionId := ctx.Param("permission_id")

	if len(strings.SplitN(permissionId, "-", 2)) != 2 {
		craneerr := cranerror.NewError(auth.CodeAccountRevokeServicePermissionParamError, "permission invalid")
		httpresponse.Error(ctx, craneerr)
		return
	}

	labels := auth.PermissionRevokeLabelKeysFromPermissionId(permissionId)

	err := a.CraneDockerClient.ServiceRemoveLabel(ctx.Param("service_id"), labels)
	if err != nil {
		craneerr := cranerror.NewError(auth.CodeAccountRevokeServicePermissionFailedError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}
	httpresponse.Ok(ctx, "success")
}
