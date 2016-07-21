package middlewares

import (
	"net/http"
	"strconv"

	"github.com/Dataman-Cloud/rolex/plugins/account"

	"github.com/gin-gonic/gin"
)

func Authorization(a *account.AccountApi) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header.Get("Authorization")) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "401"})
			ctx.Abort()
			return
		}

		value, err := a.TokenStore.Get(ctx.Request.Header.Get("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "token doesn't exists or expired"})
			ctx.Abort()
			return
		}

		accountId, _ := strconv.Atoi(value)

		acc, err := a.Authenticator.Account(uint64(accountId))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "account not found"})
			ctx.Abort()
			return
		}
		ctx.Set("account", account.ReferenceToValue(acc))

		ctx.Next()
	}
}
