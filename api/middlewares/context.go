package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func RolexApiContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rolexContext context.Context
		backgroundContext := context.Background()

		if len(ctx.Param("node_id")) > 0 {
			rolexContext = context.WithValue(backgroundContext, "node_id", ctx.Param("node_id"))
		}

		ctx.Set("rolexContext", rolexContext)
	}
}
