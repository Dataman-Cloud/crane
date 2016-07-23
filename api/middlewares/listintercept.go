package middlewares

import (
	"strconv"

	"github.com/Dataman-Cloud/rolex/model"

	"github.com/gin-gonic/gin"
)

func ListIntercept() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "GET" {
			listOptions := model.ListOptions{}
			params := ctx.Request.URL.Query()

			if p, err := strconv.ParseUint(params.Get("page"), 10, 64); err != nil {
				listOptions.Limit = 10
			} else {
				listOptions.Limit = p
			}

			if per, err := strconv.ParseUint(params.Get("per_page"), 10, 64); err != nil {
				listOptions.Offset = 0
			} else {
				listOptions.Offset = (per - 1) * listOptions.Limit
			}

			ctx.Set("listOptions", listOptions)
		}
	}
}
