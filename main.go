package main

import (
	"flag"
	"net/http"

	"github.com/Dataman-Cloud/rolex/api"
	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/Dataman-Cloud/rolex/util/config"

	"github.com/Dataman-Cloud/rolex/util/db"
	log "github.com/Dataman-Cloud/rolex/util/log"

	"golang.org/x/net/context"
)

var (
	envFile = flag.String("config", "env_file", "")
)

func main() {
	flag.Parse()
	db.InitDB()

	api := &api.Api{}

	ctx := context.Background()
	conf := config.InitConfig(*envFile)

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

	err = server.ListenAndServe()
	if err != nil {
		log.G(ctx).Fatal("can't start server: ", err)
	}
}
