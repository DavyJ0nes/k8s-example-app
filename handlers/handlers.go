package handlers

import (
	"github.com/gorilla/mux"
)

// Router is the mux Router for the Service
func Router(buildTime, commit, release string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home(buildTime, commit, release)).Methods("GET")
	return r
}
