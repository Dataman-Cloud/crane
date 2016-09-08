package authenticators

import (
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
)

type Ldap struct {
	auth.Authenticator
}

func (d *Ldap) ModificationAllowed() bool {
	return false
}
