package account

type Authenticator interface {
	Accounts(filter string) []*Account
	Account(id string) *Account
	Login(account *Account) bool
	Acls(account *Account) []*ACL
	Roles(account *Account) []*Role

	ModificationAllowed() bool
}
