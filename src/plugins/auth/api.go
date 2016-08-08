package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (a *AccountApi) CreateAccount(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if acc.Password == "" {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateParamError, "password can not be null")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if acc.Email == "" {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateParamError, "email can not be null")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateParamError, "invalid groupid")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	acc.LoginAt = time.Now()
	if err := a.Authenticator.CreateAccount(groupId, &acc); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateAuthenticatorError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GetAccountInfo(ctx *gin.Context) {
	account, _ := ctx.Get("account")
	account, err := a.Authenticator.Account(account.(Account).ID)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGetAccountNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, account)
	}
}

func (a *AccountApi) GetAccount(ctx *gin.Context) {
	account, err := a.Authenticator.Account(ctx.Param("account_id"))
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGetAccountNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, account)
	}
}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")

	accounts, err := a.Authenticator.Accounts(listOptions.(model.ListOptions))
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGetAccountNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, accounts)
	}
}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLoginParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	token, err := a.Authenticator.Login(&acc)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLoginFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	a.TokenStore.Set(ctx, token, fmt.Sprintf("%d", acc.ID), time.Now().Add(SESSION_DURATION))
	acc.Password = ""
	rolexgin.HttpOkResponse(ctx, token)
}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {
	if err := a.TokenStore.Del(ctx, ctx.Request.Header.Get("Authorization")); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLogoutError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GroupAccounts(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64); err != nil {
		log.Errorf("invalid groupid: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGroupAccountsGroupIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"group_id": groupId,
		}
	}

	accounts, err := a.Authenticator.GroupAccounts(listOptions)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGroupAccountsNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, accounts)
	}
}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAccoutGroupsAccountIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"account_id": accountId,
		}
	}

	groups, err := a.Authenticator.AccountGroups(listOptions)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAccoutGroupsNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, groups)
	}
}

func (a *AccountApi) GetGroup(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGetGroupGroupIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	group, err := a.Authenticator.Group(groupId)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGetGroupGroupIdNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, group)
	}
}

func (a *AccountApi) ListGroups(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")
	groups, err := a.Authenticator.Groups(listOptions.(model.ListOptions))
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountListGroupNotFoundError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		rolexgin.HttpOkResponse(ctx, groups)
	}
}

func (a *AccountApi) CreateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	var group Group
	if err := ctx.BindJSON(&group); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	account, _ := ctx.Get("account")
	group.CreaterId = account.(Account).ID
	if err := a.Authenticator.CreateGroup(&group); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateGroupFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) UpdateGroup(ctx *gin.Context) {
	var group Group

	if !a.Authenticator.ModificationAllowed() {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := ctx.BindJSON(&group); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountCreateGroupParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.UpdateGroup(&group); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountUpdateGroupParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) DeleteGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountDeleteGroupGroupIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
	}

	if err := a.Authenticator.DeleteGroup(groupId); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountDeleteGroupFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	rolexgin.HttpCreateResponse(ctx, "success")
}

func (a *AccountApi) JoinGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountJoinGroupGroupIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountJoinGroupAccountIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.JoinGroup(accountId, groupId); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountJoinGroupFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) LeaveGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLeaveGroupAccountIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLeaveGroupGroupIdNotValidError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.LeaveGroup(accountId, groupId); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountLeaveGroupFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return

	}

	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GrantServicePermission(ctx *gin.Context) {
	var param struct {
		GroupID uint64 `json:"GroupID"`
		Perm    string `json:"Perm"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGrantServicePermissionParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	err := a.RolexDockerClient.ServiceAddLabel(ctx.Param("service_id"), PermissionGrantLabelsPairFromGroupIdAndPerm(param.GroupID, param.Perm))
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountGrantServicePermissionFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) RevokeServicePermission(ctx *gin.Context) {
	permission_id := ctx.Param("permission_id")

	if len(strings.SplitN(permission_id, "-", 2)) != 2 {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountRevokeServicePermissionParamError, "permission invalid")
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	labels := PermissionRevokeLabelKeysFromPermissionId(permission_id)

	err := a.RolexDockerClient.ServiceRemoveLabel(ctx.Param("service_id"), labels)
	if err != nil {
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeAccountRevokeServicePermissionFailedError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	rolexgin.HttpOkResponse(ctx, "success")
}
