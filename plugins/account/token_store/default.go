package token_store

import (
	"time"

	"github.com/Dataman-Cloud/rolex/plugins/account"
)

type tokenStore struct {
	Token    string
	ExpireAt time.Time
}

type Default struct {
	account.TokenStore

	Store map[string]*tokenStore
}

func NewDefaultStore() *Default {
	return &Default{
		Store: make(map[string]*tokenStore),
	}
}

func (d *Default) Set(key, token string, expired_at time.Time) error {
	d.Store[key] = &tokenStore{Token: token, ExpireAt: expired_at}
	return nil
}

func (d *Default) Get(key string) (string, error) {
	if tokenStore, ok := d.Store[key]; ok {
		if tokenStore.ExpireAt.Before(time.Now()) {
			return tokenStore.Token, nil
		} else {
			return "", TokenExpired
		}
	} else {
		return "", TokenNotFound
	}
}
