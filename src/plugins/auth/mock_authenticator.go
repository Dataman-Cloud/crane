package auth

import (
	"github.com/Dataman-Cloud/crane/src/utils/model"
)

var (
	AuthError          error
	AccountError       error
	CreateAccountError error
	AccountsError      error
)

type MockAuthenticator struct {
	Authenticator
}

func NewMockAuthenticator() *MockAuthenticator {
	return &MockAuthenticator{}
}

func (d *MockAuthenticator) AccountPermissions(account *Account) (*[]string, error) {
	return nil, AuthError
}

func (d *MockAuthenticator) Login(a *Account) (token string, err error) {
	return "", AuthError
}

func (d *MockAuthenticator) EncryptPassword(password string) string {
	return password
}

func (d *MockAuthenticator) DeleteGroup(groupId uint64) error {
	return AuthError
}

func (d *MockAuthenticator) Groups(options model.ListOptions) (auths *[]Group, err error) {
	return nil, AuthError
}

func (d *MockAuthenticator) Group(id uint64) (*Group, error) {
	return nil, AuthError
}

func (d *MockAuthenticator) CreateGroup(g *Group) error {
	return AuthError
}

func (d *MockAuthenticator) UpdateGroup(g *Group) error {
	return AuthError
}

func (d *MockAuthenticator) GroupAccounts(account model.ListOptions) (*[]Account, error) {
	return nil, AuthError
}

func (d *MockAuthenticator) AccountGroups(account model.ListOptions) (*[]Group, error) {
	return nil, AuthError
}

func (d *MockAuthenticator) Accounts(listOptions model.ListOptions) (auths *[]Account, err error) {
	return nil, AccountsError
}

func (d *MockAuthenticator) Account(id interface{}) (*Account, error) {
	return nil, AccountError
}

func (d *MockAuthenticator) UpdateAccount(a *Account) error {
	return AuthError
}

func (d *MockAuthenticator) CreateAccount(groupId uint64, a *Account) error {
	return CreateAccountError
}

func (d *MockAuthenticator) JoinGroup(accountId, groupId uint64) error {
	return AuthError
}

func (d *MockAuthenticator) LeaveGroup(accountId, groupId uint64) error {
	return AuthError
}

func (d *MockAuthenticator) ModificationAllowed() bool {
	return true
}

func (d *MockAuthenticator) GetDefaultAccounts() (account []Account) {
	return
}
