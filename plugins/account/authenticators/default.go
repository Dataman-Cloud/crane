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

	Groups = []*account.Group{
		{ID: "1", Name: "developers"},
		{ID: "2", Name: "operation"},
	}
)

func (d *Default) ModificationAllowed() bool {
	return false
}

func (d *Default) Login(a *account.Account) (token string, err error) {
	for _, acc := range Accounts {
		if a.Password == acc.Password && a.Email == acc.Email {
			a.LoginAt = time.Now()
			a.ID = acc.ID
			return account.GenToken(a), nil
		}
	}

	return "", account.ErrLoginFailed
}

func (d *Default) Accounts(filter account.AccountFilter) (accounts []*account.Account, err error) {
	return Accounts, nil
}

func (d *Default) Account(idOrEmail string) (*account.Account, error) {
	for _, acc := range Accounts {
		if idOrEmail == acc.Email || idOrEmail == acc.ID {
			return acc, nil
		}
	}

	return nil, account.ErrAccountNotFound
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

func (d *Default) Groups(filter account.GroupFilter) (accounts []*account.Group, err error) {
	return Groups, nil
}

func (d *Default) Group(id string) (*account.Group, error) {
	for _, group := range Groups {
		if id == group.ID {
			return group, nil
		}
	}

	return nil, account.ErrGroupNotFound
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
