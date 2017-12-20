package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/davyj0nes/k8s-demo/handlers"
	"github.com/davyj0nes/k8s-demo/version"
)

func main() {
	// CONFIGURING COMMAND LINE ARGS
	port := flag.String("port", "8080", "port number to use with Web Service")
	flag.Parse()

	log.Printf("Starting Service...\ncommit: %s, build_time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("The service is ready to go...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), router))
}
