package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	var config Config

	err := config.ParseEnv()
	if err != nil {
		log.Fatalf("error occured initializing config")
	}

	log.Printf("listening port: %s\n", config.ListenPort)

	//start router
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.ListenPort, router))
}
