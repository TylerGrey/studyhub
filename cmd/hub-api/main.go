package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TylerGrey/hub-api/api/handler"
	"github.com/TylerGrey/hub-api/api/resolvers"
	"github.com/TylerGrey/hub-api/api/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
)

func main() {
	var (
		addr = flag.String("http", ":8080", "HTTP port")
	)
	flag.Parse()

	h := &handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), &resolvers.Resolver{}),
	}

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	op := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Access-Control-Allow-Headers", "DeviceInfo", "Authorization", "X-Requested-With"},
		AllowCredentials: false,
	}
	handler := cors.New(op).Handler(mux)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errc <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("Listening for requests on %s ", *addr)

	go func() {
		errc <- http.ListenAndServe(*addr, handler)
	}()

	log.Println("exit", <-errc)
}
