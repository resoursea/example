package main

import (
	"github.com/resoursea/api"
	"log"
	"net/http"
)

var route api.Router

func init() {
	// Create a Route tree to access the created Resource tree
	var err error
	route, err = api.NewRoute(Api{
		Version: 1,
		Message: "This is the REST API for a book store",
		DB:      db,
	})
	if err != nil {
		log.Fatalf("Error creating the Route: %s\n", err)
	}

	// Print the Router
	// See if the Resource tree corresponds with what it should be
	api.PrintRouter(route)
}

func main() {
	// Starting de HTTP server
	log.Println("Starting HTTP server on http://localhost:8080/")
	// The Route implements Handler interface
	// So it can be used with the standard net/http library
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}
}
