package account

type Authenticator interface {
	AccountPermissions(account *Account) []*string
	AccountGroups(account *Account) []*Group

	Login(account *Account) (token string, err error)

	Groups(filter GroupFilter) []*Group
	Group(id string) *Group
	CreateGroup(role *Group) error
	UpdateGroup(role *Group) error

	Accounts(filter AccountFilter) []*Account
	Account(id string) *Account
	CreateAccount(a *Account) error
	UpdateAccount(a *Account) error

	JoinGroup(g *Group, a *Account) error
	LeaveGroup(g *Group, a *Account) error

	ModificationAllowed() bool
}
