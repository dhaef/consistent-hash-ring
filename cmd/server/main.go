package main

import (
	"consistent-hash-ring/internal/api"
	"log"
	"net/http"
)

func main() {
	api := api.NewAPI(10)
	err := api.Setup()
	if err != nil {
		log.Fatal("failed to start api", err.Error())
	}

	server := http.Server{
		Addr:    ":4000",
		Handler: api.RegisterRoutes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server", err.Error())
	}
}
