package account

import (
	"net/url"
	"time"
)

type Account struct {
	ID       string    `json:"Id"`
	Title    string    `json:"Title"`
	Email    string    `json:"Email"`
	Phone    string    `json:"Phone"`
	LoginAt  time.Time `json:"LoginAt"`
	Password string    `json:"Password"`
	Token    string    `json:"-"`
}

type AccountFilter url.Values

func ReferenceToValue(a *Account) Account {
	return Account{
		ID:       a.ID,
		Title:    a.Title,
		Email:    a.Email,
		Phone:    a.Phone,
		LoginAt:  a.LoginAt,
		Password: a.Password,
		Token:    a.Token,
	}
}
