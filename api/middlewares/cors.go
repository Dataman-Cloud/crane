package middlewares

import (
	"github.com/gin-gonic/gin"
)

func OptionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, X-Requested-By, If-Modified-Since, X-File-Name, Cache-Control, X-XSRFToken, Authorization")
		ctx.Header("Content-Type", "application/json")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
		}

		ctx.Next()
	}
}
