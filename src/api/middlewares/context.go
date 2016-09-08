package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func CraneApiContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var craneContext context.Context
		backgroundContext := context.Background()

		if len(ctx.Param("node_id")) > 0 {
			craneContext = context.WithValue(backgroundContext, "node_id", ctx.Param("node_id"))
		}

		ctx.Set("craneContext", craneContext)
	}
}
