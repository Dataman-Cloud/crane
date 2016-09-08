package middlewares

import (
	"github.com/Dataman-Cloud/crane/src/plugins/auth"

	"github.com/gin-gonic/gin"
)

func AuthorizeServiceAccess(p auth.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
