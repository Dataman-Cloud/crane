package account

import (
	"errors"
)

var (
	ErrLoginFailed = errors.New("account login failed")
)

type Account struct {
	ID    string `json:"Id"`
	Title string `json:"Title"`
	Email string `json:"Email"`
	Phone string `json:"Phone"`
}

func NewAccount() *Account {
	a := &Account{}
	return a
}

func (account *Account) Login() (token string, err error) {
	return "", nil
}

func (account *Account) Logout() bool {
	return false
}
