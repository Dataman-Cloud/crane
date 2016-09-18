package api

import (
	"os/exec"

	containertty "github.com/Dataman-Cloud/crane/src/plugins/tty"

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
	nodeUrl, err := api.GetDockerClient().GetDaemonUrlById(ctx.Param("node_id"))
	if err != nil {
		log.Error("Get container endpoint got error: ", err)
		return
	}

	nodeUrl.Scheme = "tcp"
	cmd := exec.Command("docker", "-H", nodeUrl.String(), "exec", "-it", ctx.Param("container_id"), "sh")
	client, err := containertty.New(cmd, conn, req, containertty.DefaultOptions)
	if err != nil {
		log.Error("Create tty client got error: ", err)
		return
	}

	client.HandleClient()
	return
}
