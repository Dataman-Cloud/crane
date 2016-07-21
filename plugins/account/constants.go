package account

import (
	"errors"
	"time"
)

var (
	ErrLoginFailed     = errors.New("account login failed")
	ErrAccountNotFound = errors.New("account not found")
	ErrGroupNotFound   = errors.New("group not found")
)

const (
	SESSION_KEY_FORMAT = "account_id:%s:token"
	SESSION_DURATION   = time.Second * 60 * 10
)
