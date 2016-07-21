package authenticators

import (
	"github.com/Dataman-Cloud/rolex/plugins/account"
)

type Ldap struct {
	account.Authenticator
}

func (d *Ldap) ModificationAllowed() bool {
	return false
}
