package main

import (
	"flag"
	"log"

	"github.com/TylerGrey/study-hub/api/app"
	"github.com/joho/godotenv"
)

var addr *string

func init() {
	err := godotenv.Load("../../configs/.env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr = flag.String("http", ":8080", "HTTP server port")
	flag.Parse()
}

func main() {
	s := app.Server{
		Addr: addr,
	}
	if err := s.Start(); err != nil {
		log.Println("exit", err.Error())
	}
}
