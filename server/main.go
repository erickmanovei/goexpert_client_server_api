package main

import (
	"log"
	"net/http"

	"github.com/erickmanovei/client_server_api_server/middlewares"
	"github.com/erickmanovei/client_server_api_server/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Routes(mux)
	if err := http.ListenAndServe(":8080", middlewares.RecoverMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
