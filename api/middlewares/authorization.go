package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	ctx.Next() // do nothing with account feature
}
