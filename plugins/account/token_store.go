package account

import (
	"time"
)

type TokenStore interface {
	Set(token string, expired_at time.Time) bool
	Get(token string) bool
}
