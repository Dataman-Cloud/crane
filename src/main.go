package main

import (
	"flag"
	"net/http"

	"github.com/Dataman-Cloud/crane/src/api"
	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/plugins"
	"github.com/Dataman-Cloud/crane/src/utils/config"
	log "github.com/Dataman-Cloud/crane/src/utils/log"

	"golang.org/x/net/context"
)

func main() {
	flag.Parse()

	ctx := context.Background()
	conf := config.GetConfig()

	plugins.Init(conf)

	client, err := dockerclient.NewCraneDockerClient(conf)
	if err != nil {
		log.G(ctx).Fatal(err)
	}

	api := &api.Api{
		Client: client,
		Config: conf,
	}

	ctx = log.WithLogger(ctx, log.G(ctx).WithField("module", "main"))

	server := &http.Server{
		Addr:           conf.CraneAddr,
		Handler:        api.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.G(ctx).Fatal("can't start server: ", err)
	}
}
