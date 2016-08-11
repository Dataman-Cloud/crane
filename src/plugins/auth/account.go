package auth

import (
	"net/url"
	"time"
)

type Account struct {
	ID       uint64    `json:"Id"`
	Title    string    `json:"Title"`
	Email    string    `json:"Email" gorm:"not null"`
	Phone    string    `json:"Phone"`
	LoginAt  time.Time `json:"LoginAt"`
	Password string    `json:"Password" gorm:"not null"`
	Token    string    `json:"-" gorm:"-"`
}

type AccountFilter url.Values

type AccountGroup struct {
	ID        uint64 `json:"Id"`
	AccountId uint64 `json:"AccountId"`
	GroupId   uint64 `json:"GroupId"`
}

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
