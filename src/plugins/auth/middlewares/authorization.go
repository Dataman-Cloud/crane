package middlewares

import (
	"strconv"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"
	"github.com/Dataman-Cloud/crane/src/utils/model"

	"github.com/gin-gonic/gin"
)

func Authorization(tokenStore auth.TokenStore, authenticator auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Query("Authorization")) != 0 {
			ctx.Request.Header.Set("Authorization", ctx.Query("Authorization"))
		}

		if len(ctx.Query("Cookie")) != 0 {
			ctx.Request.Header.Set("Cookie", ctx.Query("Cookie"))
		}

		if len(ctx.Request.Header.Get("Authorization")) == 0 {
			httpresponse.Error(ctx, cranerror.NewError(auth.CodeAccountTokenInvalidError, "Invalid Authorization"))
			ctx.Abort()
			return
		}

		value, err := tokenStore.Get(ctx, ctx.Request.Header.Get("Authorization"))
		if err != nil {
			httpresponse.Error(ctx, cranerror.NewError(auth.CodeAccountTokenInvalidError, "Invalid Authorization"))
			ctx.Abort()
			return
		}

		accountId, _ := strconv.ParseUint(value, 10, 64)

		acc, err := authenticator.Account(accountId)
		if err != nil {
			httpresponse.Error(ctx, cranerror.NewError(auth.CodeAccountTokenInvalidError, "Invalid Authorization"))
			ctx.Abort()
			return
		}

		ctx.Set("account", auth.ReferenceToValue(acc))

		if groups, err := authenticator.AccountGroups(model.ListOptions{
			Filter: map[string]interface{}{
				"account_id": accountId,
			},
		}); err != nil {
			httpresponse.Error(ctx, cranerror.NewError(auth.CodeAccountTokenInvalidError, "Invalid Authorization"))
			ctx.Abort()
			return
		} else {
			ctx.Set("groups", *groups)
		}

		ctx.Next()
	}
}
