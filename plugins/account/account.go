package account

import (
	"errors"
	"net/url"
	"time"
)

var (
	ErrLoginFailed = errors.New("account login failed")
)

type Account struct {
	ID       string    `json:"Id"`
	Title    string    `json:"Title"`
	Email    string    `json:"Email"`
	Phone    string    `json:"Phone"`
	Password string    `json:"Password"`
	Token    string    `json:"Token"`
	LoginAt  time.Time `json:"LoginAt"`
}

type AccountFilter url.Values
