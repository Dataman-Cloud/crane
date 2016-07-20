package authenticators

import (
	"time"

	"github.com/Dataman-Cloud/rolex/plugins/account"
)

type Default struct {
	account.Authenticator
}

func NewDefaultAuthenticator() *Default {
	return &Default{}
}

var (
	Accounts = []*account.Account{
		{ID: "1", Title: "Engineering", Email: "admin@admin.com", Password: "adminadmin"},
	}
)

func (d *Default) ModificationAllowed() bool {
	return false
}

func (d *Default) Login(a *account.Account) (token string, err error) {
	for _, account := range Accounts {
		if a.Password == account.Password && a.Email == account.Email {
			a.LoginAt = time.Now()
			return "", nil
		}
	}
	return "", nil
}

func (d *Default) Accounts(filter account.AccountFilter) []*account.Account {
	return nil
}

func (d *Default) Account(idOrEmail string) *account.Account {
	return nil
}

func (d *Default) DeleteAccount(a *account.Account) error {
	return nil
}

func (d *Default) UpdateAccount(a *account.Account) error {
	return nil
}

func (d *Default) CreateAccount(a *account.Account) error {
	return nil
}

func (d *Default) Groups(filter account.GroupFilter) []*account.Group {
	return nil
}

func (d *Default) Group(id string) *account.Group {
	return nil
}

func (d *Default) CreateGroup(g *account.Group) error {
	return nil
}

func (d *Default) DeleteGroup(g *account.Group) error {
	return nil
}

func (d *Default) UpdateGroup(g *account.Group) error {
	return nil
}

func (d *Default) JoinGroup(g *account.Group, a *account.Account) error {
	return nil
}

func (d *Default) LeaveGroup(g *account.Group, a *account.Account) error {
	return nil
}
