package token_store

import (
	"errors"
)

var (
	TokenExpired  = errors.New("token expired")
	TokenNotFound = errors.New("token not found")
)
