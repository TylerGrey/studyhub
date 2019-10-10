package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
)

type Server struct {
	Addr    *string
	Handler http.Handler
}

func (s Server) Start() error {
	op := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Access-Control-Allow-Headers", "DeviceInfo", "Authorization", "X-Requested-With"},
		AllowCredentials: false,
	}
	handler := cors.New(op).Handler(s.Handler)

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
