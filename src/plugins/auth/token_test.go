package auth

import (
	"testing"
)

func TestGenToken(T *testing.T) {
	account := &Account{
		Email:    "admin@admin.com",
		Password: "admin",
	}
	if token := GenToken(account); token != "549aca403f9ef5d072f407f7620f3f86" {
		T.Error(token)
	}
}
