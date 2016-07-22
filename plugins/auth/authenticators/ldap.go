package authenticators

import (
	"github.com/Dataman-Cloud/rolex/plugins/auth"
)

type Ldap struct {
	auth.Authenticator
}

func (d *Ldap) ModificationAllowed() bool {
	return false
}
