package api

import (
	"os/exec"

	"github.com/Dataman-Cloud/rolex/plugins/tty"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (api *Api) ConnectContainer(ctx *gin.Context) {
	req := ctx.Request
	conn, err := containertty.Upgrader.Upgrade(ctx.Writer, req, nil)
	if err != nil {
		log.Error("Upgrade websocket connect got error: ", err)
		return
	}

	_, stream, err := conn.ReadMessage()
	if err != nil {
		log.Error("Get websocket init message got error: ", err)
		return
	}
	log.Info("Init message: ", string(stream))

	//cmd := exec.Command("ping", "www.baidu.com")
	cmd := exec.Command("ssh", "root@192.168.59.103")
	client, err := containertty.New(cmd, conn, req, containertty.DefaultOptions)
	if err != nil {
		log.Error("Create tty client got error: ", err)
		return
	}

	client.HandleClient()
	return
}
