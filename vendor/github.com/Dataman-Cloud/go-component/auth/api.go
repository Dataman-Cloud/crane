package auth

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/go-component/utils/dmgin"
	"github.com/Dataman-Cloud/go-component/utils/model"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	//Account
	CodeInvalidGroupId                                  = "400-12001"
	CodeAccountCreateParamError                         = "400-12002"
	CodeAccountCreateAuthenticatorError                 = "503-12003"
	CodeAccountGetAccountError                          = "503-12004"
	CodeAccountGetAccountNotFoundError                  = "503-12005"
	CodeAccountLoginParamError                          = "400-12006"
	CodeAccountLoginFailedError                         = "401-12007"
	CodeAccountLogoutError                              = "503-12008"
	CodeAccountGroupAccountsGroupIdNotValidError        = "400-12009"
	CodeAccountGroupAccountsNotFoundError               = "404-12010"
	CodeAccountAccoutGroupsAccountIdNotValidError       = "400-12011"
	CodeAccountAccoutGroupsNotFoundError                = "503-12012"
	CodeAccountGetGroupGroupIdNotValidError             = "503-12013"
	CodeAccountGetGroupGroupIdNotFoundError             = "404-12014"
	CodeAccountListGroupNotFoundError                   = "503-12015"
	CodeAccountAuthenticatorModificationNotAllowedError = "503-12016"
	CodeAccountCreateGroupParamError                    = "400-12017"
	CodeAccountCreateGroupFailedError                   = "503-12018"
	CodeAccountUpdateGroupParamError                    = "400-12019"
	CodeAccountUpdateGroupFailedError                   = "503-12020"
	CodeAccountDeleteGroupGroupIdNotValidError          = "503-12021"
	CodeAccountDeleteGroupFailedError                   = "503-12022"
	CodeAccountJoinGroupGroupIdNotValidError            = "503-12023"
	CodeAccountJoinGroupAccountIdNotValidError          = "503-12024"
	CodeAccountJoinGroupFailedError                     = "503-12025"
	CodeAccountLeaveGroupGroupIdNotValidError           = "503-12026"
	CodeAccountLeaveGroupAccountIdNotValidError         = "503-12027"
	CodeAccountLeaveGroupFailedError                    = "503-12028"
	CodeAccountGrantServicePermissionParamError         = "400-12029"
	CodeAccountGrantServicePermissionFailedError        = "503-12030"
	CodeAccountRevokeServicePermissionParamError        = "400-12031"
	CodeAccountRevokeServicePermissionFailedError       = "503-12032"
	CodeAccountTokenInvalidError                        = "401-12033"
	CodeAccountLoginFailedEmailNotValidError            = "401-12034"
	CodeAccountLoginFailedPasswordNotValidError         = "401-12035"
)

func (a *AccountApi) CreateAccount(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if acc.Password == "" {
		rolexerr := dmerror.NewError(CodeAccountCreateParamError, "password can not be null")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if acc.Email == "" {
		rolexerr := dmerror.NewError(CodeAccountCreateParamError, "email can not be null")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateParamError, "invalid groupid")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	acc.Password = a.Authenticator.EncryptPassword(acc.Password)
	acc.LoginAt = time.Now()
	if err := a.Authenticator.CreateAccount(groupId, &acc); err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateAuthenticatorError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GetAccountInfo(ctx *gin.Context) {
	account, _ := ctx.Get("account")
	account, err := a.Authenticator.Account(account.(Account).ID)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGetAccountNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, account)
	}
}

func (a *AccountApi) GetAccount(ctx *gin.Context) {
	account, err := a.Authenticator.Account(ctx.Param("account_id"))
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGetAccountNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, account)
	}
}

func (a *AccountApi) ListAccounts(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")

	accounts, err := a.Authenticator.Accounts(listOptions.(model.ListOptions))
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGetAccountNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, accounts)
	}
}

func (a *AccountApi) AccountLogin(ctx *gin.Context) {
	var acc Account
	if err := ctx.BindJSON(&acc); err != nil {
		rolexerr := dmerror.NewError(CodeAccountLoginParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	token, err := a.Authenticator.Login(&acc)
	if err != nil {
		dmgin.HttpErrorResponse(ctx, err)
		return
	}

	a.TokenStore.Set(ctx, token, fmt.Sprintf("%d", acc.ID), time.Now().Add(SESSION_DURATION))
	acc.Password = ""
	dmgin.HttpOkResponse(ctx, token)
}

func (a *AccountApi) AccountLogout(ctx *gin.Context) {
	if err := a.TokenStore.Del(ctx, ctx.Request.Header.Get("Authorization")); err != nil {
		rolexerr := dmerror.NewError(CodeAccountLogoutError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GroupAccounts(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64); err != nil {
		log.Errorf("invalid groupid: %v", err)
		rolexerr := dmerror.NewError(CodeAccountGroupAccountsGroupIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"group_id": groupId,
		}
	}

	accounts, err := a.Authenticator.GroupAccounts(listOptions)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGroupAccountsNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, accounts)
	}
}

func (a *AccountApi) AccountGroups(ctx *gin.Context) {
	listObj, _ := ctx.Get("listOptions")
	listOptions := listObj.(model.ListOptions)

	if accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64); err != nil {
		rolexerr := dmerror.NewError(CodeAccountAccoutGroupsAccountIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		listOptions.Filter = map[string]interface{}{
			"account_id": accountId,
		}
	}

	groups, err := a.Authenticator.AccountGroups(listOptions)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountAccoutGroupsNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, groups)
	}
}

func (a *AccountApi) GetGroup(ctx *gin.Context) {
	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGetGroupGroupIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	group, err := a.Authenticator.Group(groupId)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGetGroupGroupIdNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, group)
	}
}

func (a *AccountApi) ListGroups(ctx *gin.Context) {
	listOptions, _ := ctx.Get("listOptions")
	groups, err := a.Authenticator.Groups(listOptions.(model.ListOptions))
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountListGroupNotFoundError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	} else {
		dmgin.HttpOkResponse(ctx, groups)
	}
}

func (a *AccountApi) CreateGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := dmerror.NewError(CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	var group Group
	if err := ctx.BindJSON(&group); err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	account, _ := ctx.Get("account")
	group.CreaterId = account.(Account).ID
	if err := a.Authenticator.CreateGroup(&group); err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateGroupFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) UpdateGroup(ctx *gin.Context) {
	var group Group

	if !a.Authenticator.ModificationAllowed() {
		rolexerr := dmerror.NewError(CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := ctx.BindJSON(&group); err != nil {
		rolexerr := dmerror.NewError(CodeAccountCreateGroupParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.UpdateGroup(&group); err != nil {
		rolexerr := dmerror.NewError(CodeAccountUpdateGroupParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) DeleteGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := dmerror.NewError(CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountDeleteGroupGroupIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
	}

	if err := a.Authenticator.DeleteGroup(groupId); err != nil {
		rolexerr := dmerror.NewError(CodeAccountDeleteGroupFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpCreateResponse(ctx, "success")
}

func (a *AccountApi) JoinGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := dmerror.NewError(CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountJoinGroupGroupIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountJoinGroupAccountIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.JoinGroup(accountId, groupId); err != nil {
		rolexerr := dmerror.NewError(CodeAccountJoinGroupFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) LeaveGroup(ctx *gin.Context) {
	if !a.Authenticator.ModificationAllowed() {
		rolexerr := dmerror.NewError(CodeAccountAuthenticatorModificationNotAllowedError, "moditication not allowed")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountLeaveGroupAccountIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	groupId, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountLeaveGroupGroupIdNotValidError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if err := a.Authenticator.LeaveGroup(accountId, groupId); err != nil {
		rolexerr := dmerror.NewError(CodeAccountLeaveGroupFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return

	}

	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) GrantServicePermission(ctx *gin.Context) {
	var param struct {
		GroupID uint64 `json:"GroupID"`
		Perm    string `json:"Perm"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		rolexerr := dmerror.NewError(CodeAccountGrantServicePermissionParamError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	err := a.RolexDockerClient.ServiceAddLabel(ctx.Param("service_id"), PermissionGrantLabelsPairFromGroupIdAndPerm(param.GroupID, param.Perm))
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountGrantServicePermissionFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}

func (a *AccountApi) RevokeServicePermission(ctx *gin.Context) {
	permission_id := ctx.Param("permission_id")

	if len(strings.SplitN(permission_id, "-", 2)) != 2 {
		rolexerr := dmerror.NewError(CodeAccountRevokeServicePermissionParamError, "permission invalid")
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	labels := PermissionRevokeLabelKeysFromPermissionId(permission_id)

	err := a.RolexDockerClient.ServiceRemoveLabel(ctx.Param("service_id"), labels)
	if err != nil {
		rolexerr := dmerror.NewError(CodeAccountRevokeServicePermissionFailedError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}
	dmgin.HttpOkResponse(ctx, "success")
}
