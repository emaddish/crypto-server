package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	var config Config

	config.ListenPort = "10000"
	config.CryptoURL = "https://api.hitbtc.com/api/2/public"

	log.Printf("listening port: %s\n", config.ListenPort)

	//start router
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.ListenPort, router))
}
