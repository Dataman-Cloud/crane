package main

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/router"

	log "github.com/Dataman-Cloud/rolex/util/log"
	"golang.org/x/net/context"
)

func main() {

	ctx := context.Background()
	ctx = log.WithLogger(ctx, log.G(ctx).WithField("module", "main"))

	server := &http.Server{
		Addr:           "0.0.0.0:5013",
		Handler:        router.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.G(ctx).Debug("can't start server: ", err)
	}
}
