package middlewares

import (
	"github.com/Dataman-Cloud/rolex/src/plugins/auth"

	"github.com/gin-gonic/gin"
)

func AuthorizeNetworkAccess(p auth.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
