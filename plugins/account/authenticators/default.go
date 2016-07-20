package authenticators

import (
	"github.com/Dataman-Cloud/rolex/plugins/account"
)

type Default struct {
	account.Authenticator
}

func (d *Default) ModificationAllowed() bool {
	return true
}
