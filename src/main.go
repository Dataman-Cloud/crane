package main

import (
	"flag"
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/api"
	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	//"github.com/Dataman-Cloud/rolex/src/util/db"
	log "github.com/Dataman-Cloud/rolex/src/util/log"

	"golang.org/x/net/context"
)

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

	err = server.ListenAndServe()
	if err != nil {
		log.G(ctx).Fatal("can't start server: ", err)
	}
}
