package main

import (
	"github.com/Dataman-Cloud/newworld/rolex-go/router/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
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
