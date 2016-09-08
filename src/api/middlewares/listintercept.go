package middlewares

import (
	"strconv"

	"github.com/Dataman-Cloud/crane/src/utils/model"

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

			params.Del("page")
			params.Del("per_page")
			if len(params) > 0 {
				filter := make(map[string]interface{})
				for k, v := range params {
					filter[k] = v[0]
				}
				listOptions.Filter = filter
			}

			ctx.Set("listOptions", listOptions)
		}
	}
}
