package account

import (
	"time"
)

type TokenStore interface {
	Set(key, token string, expired_at time.Time) error
	Get(key string) (string, error)
	Del(key string) error
}
