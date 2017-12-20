package main

import (
	"log"
	"net/http"

	"github.com/davyj0nes/k8s-demo/handlers"
)

func main() {
	log.Print("Starting Service...")
	router := handlers.Router()
	log.Print("The service is ready to go...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
