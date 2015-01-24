package main

import (
	"github.com/resoursea/api"
	"log"
	"net/http"
)

var route *api.Route

func init() {
	// Creating a new Api Resource tree
	// Here is added the initial values for the resources
	resource, err := api.NewResource(Api{
		Version: 1,
		Message: "This is the REST API for a book store",
		DB:      db,
	})
	if err != nil {
		log.Fatalf("Error creating the Api resource: %s\n", err)
	}

	// Create a Route tree to access the created Resource tree
	route, err = api.NewRoute(resource)
	if err != nil {
		log.Fatalf("Error creating the Route: %s\n", err)
	}

	// Print the Resource and Route trees
	//api.PrintResource(resource)
	//api.PrintRoute(route)
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
