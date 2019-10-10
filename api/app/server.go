package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TylerGrey/hub-api/api/app/handler"
	"github.com/TylerGrey/hub-api/api/app/repo/user"
	"github.com/TylerGrey/hub-api/api/app/resolvers"
	"github.com/TylerGrey/hub-api/api/app/schema"
	"github.com/TylerGrey/hub-api/internal/mysql"
	"github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
)

type Server struct {
	Addr *string
}

func (s Server) Start() error {
	h := &handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), &resolvers.Resolver{}),
	}

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	userMaster, userReplica, err := mysql.IntializeDatabase(os.Getenv("MYSQL_TABLE_USER"))
	if err != nil {
		log.Println("db initialize error", err.Error())
		panic(err)
	}
	user.NewUserRepository(userMaster, userReplica)

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

	log.Printf("Listening for requests on %s", *s.Addr)

	go func() {
		errc <- http.ListenAndServe(*s.Addr, handler)
	}()

	return <-errc
}
