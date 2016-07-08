package main

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/router"

	log "github.com/Sirupsen/logrus"
)

func main() {
	server := &http.Server{
		Addr:           "0.0.0.0:5013",
		Handler:        router.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("can't start server: ", err)
	}
}
