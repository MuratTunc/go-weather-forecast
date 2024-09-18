package main

import (
	"log"
	"net/http"
)

const port = ":8089"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("WEATHER-SERVICE is starting on port %s\n", port)

	// define http server
	srv := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("WEATHER-SERVICE is started no error... on port %s\n", port)
}
