package auth

import (
	"time"

	"github.com/gin-gonic/gin"
)

type TokenStore interface {
	Set(ctx *gin.Context, token, account_id string, expired_at time.Time) error
	Get(ctx *gin.Context, token string) (string, error)
	Del(ctx *gin.Context, token string) error
}
