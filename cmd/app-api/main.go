package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/TylerGrey/hub-api/api/app"
	"github.com/TylerGrey/hub-api/api/app/handler"
	"github.com/TylerGrey/hub-api/api/app/resolvers"
	"github.com/TylerGrey/hub-api/api/app/schema"
	graphql "github.com/graph-gophers/graphql-go"
)

var addr *string

func init() {
	addr = flag.String("http", ":8080", "HTTP server port")
	flag.Parse()
}

func main() {
	h := &handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), &resolvers.Resolver{}),
	}

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	s := app.Server{
		Addr:    addr,
		Handler: mux,
	}
	if err := s.Start(); err != nil {
		log.Println("exit", err.Error())
	}
}
