package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	app := application{}

	app.Session = getSession()

	mux := app.routes()

	log.Println("Starting servr on 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server with %s", err.Error())
	}
}
