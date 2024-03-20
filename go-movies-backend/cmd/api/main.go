package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// TODO: set up application config
	var app application

	// TODO: read from command line

	// TODO: connect to the database

	app.Domain = "example.com"
	log.Println("Starting the application on port", port)

	// TODO: start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
