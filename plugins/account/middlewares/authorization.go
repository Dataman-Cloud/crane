package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Dataman-Cloud/rolex/plugins/account"
)

func Authorization(a *account.AccountApi) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Get("Authorization")) == 0 || len(ctx.Request.Header.Get("AccountId")) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "401"})
			ctx.Abort()
			return
		}

		token, err := a.TokenStore.Get(fmt.Sprintf(account.SESSION_KEY_FORMAT, ctx.Request.Header.Get("AccountId")))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "token doesn't exists or expired"})
			ctx.Abort()
			return
		}

		if ctx.Request.Header.Get("Authorization") != token {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "token invalid"})
			ctx.Abort()
			return
		}

		acc, err := a.Authenticator.Account(ctx.Request.Header.Get("AccountId"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "account not found"})
			ctx.Abort()
			return
		}
		ctx.Set("account", account.ReferenceToValue(acc))

		ctx.Next()
	}
}
