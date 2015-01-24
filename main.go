package main

import (
	"flag"
	"fmt"
	"github.com/resourcerest/api"
	"log"
	"net/http"
)

var route *api.Route

// Development Env
var env = Env{
	Url:        "http://localhost:8080/",
	Port:       8080,
	Production: false,
}
var prod bool

func init() {

	flag.BoolVar(&prod, "prod", false, "Production? True or False.")

	resource, err := api.NewResource(Api{
		Version: 0,
		DB:      db,
		Env:     env,
	})
	if err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}

	route, err = api.NewRoute(resource)
	if err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}

	// Print TESTS
	//api.PrintResource(resource)
	//api.PrintRoute(route)
}

func main() {
	flag.Parse()
	if prod {
		env = Env{
			Url:        "http://localhost:8080/",
			Port:       8080,
			Production: true,
		}
	}

	// Starting de HTTP server
	log.Println("Starting HTTP server in " + env.Url + " ...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.Port), route); err != nil {
		log.Fatalf("Server Fatal: %s\n", err)
	}
}
