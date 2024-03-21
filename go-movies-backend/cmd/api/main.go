package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/huannguyen2114/Another-Golang-Project/internal/repository"
	"github.com/huannguyen2114/Another-Golang-Project/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	// TODO: set up application config
	var app application

	// TODO: read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// TODO: connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	defer app.DB.Connection().Close()

	app.Domain = "example.com"
	log.Println("Starting the application on port", port)

	// TODO: start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
