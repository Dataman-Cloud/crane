package auth

import (
	"github.com/Dataman-Cloud/crane/src/utils/model"
)

type Authenticator interface {
	AccountPermissions(account *Account) (*[]string, error)

	Login(account *Account) (token string, err error)
	EncryptPassword(password string) string

	DeleteGroup(groupId uint64) error
	Groups(listOptions model.ListOptions) (*[]Group, error)
	Group(id uint64) (*Group, error)
	CreateGroup(role *Group) error
	UpdateGroup(role *Group) error
	GroupAccounts(account model.ListOptions) (*[]Account, error)

	AccountGroups(account model.ListOptions) (*[]Group, error)
	Accounts(listOptions model.ListOptions) (*[]Account, error)
	Account(id interface{}) (*Account, error)
	CreateAccount(groupId uint64, a *Account) error
	UpdateAccount(a *Account) error

	JoinGroup(accountId, groupId uint64) error
	LeaveGroup(accountId, groupId uint64) error

	ModificationAllowed() bool
	GetDefaultAccounts() []Account
}
