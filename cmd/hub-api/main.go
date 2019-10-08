package main

import (
	"log"
	"net/http"

	"github.com/TylerGrey/hub-api/api/resolvers"
	"github.com/TylerGrey/hub-api/api/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schema := graphql.MustParseSchema(schema.GetRootSchema(), &resolvers.Resolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
