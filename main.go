package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davyj0nes/k8s-demo/handlers"
	"github.com/davyj0nes/k8s-demo/version"
)

func main() {
	// CONFIGURING COMMAND LINE ARGS
	port := flag.String("port", "8080", "port number to use with Web Service")
	flag.Parse()

	log.Printf("Starting Service... | {commit: %s, build_time: %s, release: %s}", version.Commit, version.BuildTime, version.Release)
	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: router,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Print("The service is ready to go...")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
