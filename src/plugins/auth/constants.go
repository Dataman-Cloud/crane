package auth

import (
	"errors"
	"time"
)

var (
	ErrLoginFailed     = errors.New("account login failed")
	ErrAccountNotFound = errors.New("account not found")
	ErrGroupNotFound   = errors.New("group not found")
)

const (
	SESSION_DURATION   = time.Second * 60 * 10 * 1000
	SESSION_KEY_FORMAT = "account_id:%v:token"
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
