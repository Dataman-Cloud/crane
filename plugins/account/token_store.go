package account

import (
	"time"
)

type TokenStore interface {
	Set(token, account_id string, expired_at time.Time) error
	Get(token string) (string, error)
	Del(token string) error
}
