package authenticators

import (
	"github.com/Dataman-Cloud/go-component/auth"
)

type Ldap struct {
	auth.Authenticator
}

func (d *Ldap) ModificationAllowed() bool {
	return false
}
