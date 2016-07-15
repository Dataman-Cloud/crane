package api

import (
	"io"

	"github.com/gin-gonic/gin"
)

func (api *Api) Logs(ctx *gin.Context) {
	message := make(chan string)

	defer close(message)

	go api.GetDockerClient().Logs(ctx.Param("node_id"), ctx.Param("container_id"), message)

	ctx.Stream(func(w io.Writer) bool {
		w.Write([]byte(<-message))
		return true
	})
}
