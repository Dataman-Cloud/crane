package api

import (
	"os/exec"

	"github.com/Dataman-Cloud/rolex/src/plugins/tty"

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

	// Add TLS config and file path
	endpoint, err := api.GetDockerClient().NodeDaemonEndpoint(ctx.Param("node_id"), "tcp")
	if err != nil {
		log.Error("Get container endpoint got error: ", err)
		return
	}

	cmd := exec.Command("docker", "-H", endpoint, "exec", "-it", ctx.Param("container_id"), "sh")
	client, err := containertty.New(cmd, conn, req, containertty.DefaultOptions)
	if err != nil {
		log.Error("Create tty client got error: ", err)
		return
	}

	client.HandleClient()
	return
}
