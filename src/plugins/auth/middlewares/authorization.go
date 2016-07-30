package middlewares

import (
	//"fmt"
	"net/http"
	"strconv"
	//"time"

	"github.com/Dataman-Cloud/rolex/src/plugins/auth"

	"github.com/gin-gonic/gin"
)

func Authorization(a *auth.AccountApi) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Query("Authorization")) != 0 {
			ctx.Request.Header.Set("Authorization", ctx.Query("Authorization"))
		}

		if len(ctx.Query("Cookie")) != 0 {
			ctx.Request.Header.Set("Cookie", ctx.Query("Cookie"))
		}

		if len(ctx.Request.Header.Get("Authorization")) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "401"})
			ctx.Abort()
			return
		}

		value, err := a.TokenStore.Get(ctx, ctx.Request.Header.Get("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "token doesn't exists or expired"})
			ctx.Abort()
			return
		}

		accountId, _ := strconv.ParseUint(value, 10, 64)

		acc, err := a.Authenticator.Account(accountId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "data": "account not found"})
			ctx.Abort()
			return
		}

		//a.TokenStore.Set(ctx, ctx.Request.Header.Get("Authorization"), fmt.Sprintf("%d", acc.ID), time.Now().Add(auth.SESSION_DURATION))
		ctx.Set("account", auth.ReferenceToValue(acc))

		ctx.Next()
	}
}
