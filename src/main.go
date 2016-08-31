package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"net"
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/api"
	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	//"github.com/Dataman-Cloud/rolex/src/util/db"
	log "github.com/Dataman-Cloud/rolex/src/util/log"

	"golang.org/x/net/context"
)

const ALIVE_URL = "http://localhost:4500"

var (
	envFile = flag.String("config", "env_file", "")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	conf := config.InitConfig(*envFile)

	//db.InitDB()

	client, err := dockerclient.NewRolexDockerClient(conf)
	if err != nil {
		log.G(ctx).Fatal(err)
	}

	api := &api.Api{
		Client: client,
		Config: conf,
	}

	ctx = log.WithLogger(ctx, log.G(ctx).WithField("module", "main"))

	server := &http.Server{
		Addr:           conf.RolexAddr,
		Handler:        api.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	go alive()

	err = server.ListenAndServe()
	if err != nil {
		log.G(ctx).Fatal("can't start server: ", err)
	}
}

func alive() {
	var id struct {
		UniqId string `json:"UniqId"`
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	if len(interfaces) == 0 {
		return
	}
	id.UniqId = interfaces[len(interfaces)-1].HardwareAddr.String()
	jsonStr, _ := json.Marshal(id)

	req, err := http.NewRequest("POST", ALIVE_URL+"/activities", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
