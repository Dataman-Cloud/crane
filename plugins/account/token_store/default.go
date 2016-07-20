package token_store

import (
	"fmt"
	"time"

	"github.com/Dataman-Cloud/rolex/plugins/account"
	log "github.com/Sirupsen/logrus"
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
	log.Debugf("Set ", key, " ", token, " ", expired_at)
	d.Store[key] = &tokenStore{Token: token, ExpireAt: expired_at}
	return nil
}

func (d *Default) Get(key string) (string, error) {
	log.Debugf("Get ", key)
	if tokenStore, ok := d.Store[key]; ok {
		if tokenStore.ExpireAt.After(time.Now()) {
			log.Debugf("Get ", tokenStore.Token)
			return tokenStore.Token, nil
		} else {
			fmt.Println("expird")
			return "", TokenExpired
		}
	} else {
		return "", TokenNotFound
	}
}

func (d *Default) Del(key string) error {
	delete(d.Store, key)
	return nil
}
