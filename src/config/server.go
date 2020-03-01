package config

import (
	"log"
	"net/http"

	"github.com/estebanborai/go-spacex-graphql/src/service"

	"github.com/graph-gophers/graphql-go/relay"
)

// Serve bootstraps the server and initializes it
func Serve() {
	log.Println("Initializing Server")

	services := service.InitProvider()

	schema := NewGraphQLSchema(services)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Println("Running at http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
