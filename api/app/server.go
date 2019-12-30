package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TylerGrey/studyhub/api/app/handler"
	"github.com/TylerGrey/studyhub/api/app/resolvers"
	"github.com/TylerGrey/studyhub/api/app/schema"
	mysqlLib "github.com/TylerGrey/studyhub/internal/mysql"
	"github.com/TylerGrey/studyhub/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
)

// Server API Server
type Server struct {
	Addr *string
}

// Start run server
func (s Server) Start() error {
	// DB 설정
	mysqlMaster, mysqlReplica, err := mysqlLib.IntializeDatabase(os.Getenv("MYSQL_DB_NAME"))
	if err != nil {
		log.Println("db initialize error", err.Error())
		panic(err)
	}
	userRepo := repo.NewUserRepository(mysqlMaster, mysqlReplica)
	hubRepo := repo.NewHubRepository(mysqlMaster, mysqlReplica)
	hubIncorrectInfoRepo := repo.NewHubIncorrectInfoRepository(mysqlMaster, mysqlReplica)

	// Handler 설정
	h := &handler.GraphQL{
		Schema: graphql.MustParseSchema(schema.GetRootSchema(), &resolvers.Resolver{
			UserRepo:             userRepo,
			HubRepo:              hubRepo,
			HubIncorrectInfoRepo: hubIncorrectInfoRepo,
		}),
	}

	mux := http.NewServeMux()
	mux.Handle("/", &handler.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h)

	// CORS 설정
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
