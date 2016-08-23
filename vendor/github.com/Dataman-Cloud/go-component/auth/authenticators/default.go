package authenticators

import (
	"time"

	"github.com/Dataman-Cloud/go-component/auth"
	"github.com/Dataman-Cloud/go-component/utils/model"
	"github.com/Dataman-Cloud/rolex/src/util/config"
)

type Default struct {
	auth.Authenticator
}

func NewDefaultAuthenticator() *Default {
	return &Default{}
}

var (
	Accounts = []auth.Account{
		{ID: 1, Title: "Engineering", Email: "admin@admin.com", Password: "adminadmin"},
	}

	Groups = []auth.Group{
		{ID: 1, Name: "developers"},
		{ID: 2, Name: "operation"},
	}
)

func (d *Default) GetDefaultAccounts() []auth.Account {
	if config.GetConfig().AccountEmailDefault != "" &&
		config.GetConfig().AccountPasswordDefault != "" {
		Accounts = []auth.Account{
			{
				ID:       1,
				Title:    "Engineering",
				Email:    config.GetConfig().AccountEmailDefault,
				Password: config.GetConfig().AccountPasswordDefault,
			},
		}
	}
	return Accounts
}

func (d *Default) ModificationAllowed() bool {
	return false
}

func (d *Default) Login(a *auth.Account) (token string, err error) {
	for _, acc := range d.GetDefaultAccounts() {
		if a.Password == acc.Password && a.Email == acc.Email {
			a.LoginAt = time.Now()
			a.ID = acc.ID
			return auth.GenToken(a), nil
		}
	}

	return "", auth.ErrLoginFailed
}

func (d *Default) Accounts(listOptions model.ListOptions) (auths *[]auth.Account, err error) {
	return &Accounts, nil
}

func (d *Default) Account(idOrEmail interface{}) (*auth.Account, error) {
	for _, acc := range Accounts {
		if id, ok := idOrEmail.(uint64); ok && acc.ID == id {
			return &acc, nil
		} else if email, ok := idOrEmail.(string); ok && acc.Email == email {
			return &acc, nil
		}
	}

	return nil, auth.ErrAccountNotFound
}

func (d *Default) DeleteAccount(a *auth.Account) error {
	return nil
}

func (d *Default) UpdateAccount(a *auth.Account) error {
	return nil
}

func (d *Default) CreateAccount(groupId uint64, a *auth.Account) error {
	return nil
}

func (d *Default) AccountGroups(listOptions model.ListOptions) (auths *[]auth.Group, err error) {
	return &Groups, nil
}

func (d *Default) Groups(listOptions model.ListOptions) (auths *[]auth.Group, err error) {
	return &Groups, nil
}

func (d *Default) Group(id uint64) (*auth.Group, error) {
	for _, group := range Groups {
		if id == group.ID {
			return &group, nil
		}
	}

	return nil, auth.ErrGroupNotFound
}

func (d *Default) CreateGroup(g *auth.Group) error {
	return nil
}

func (d *Default) DeleteGroup(groupId uint64) error {
	return nil
}

func (d *Default) UpdateGroup(g *auth.Group) error {
	return nil
}

func (d *Default) JoinGroup(accountId, groupId uint64) error {
	return nil
}

func (d *Default) LeaveGroup(accountId, groupId uint64) error {
	return nil
}

func (d *Default) EncryptPassword(password string) string {
	return password
}
