package main

import (
	"log"
	"net/http"

	"github.com/eventmodeling/workshop-warsaw/register/infrastructure/events"

	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/eventmodeling/workshop-warsaw/register/router"
)

const eventsDirectory = "/events"

func main() {
	publisher := events.Publisher{eventsDirectory}
	eventsReader := events.Reader{eventsDirectory}

	registerHandler := register.RegisterHandler{
		Publisher:    publisher,
		EventsReader: eventsReader,
	}

	r := router.NewRouter(registerHandler)
	log.Print("Running registration server")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalf("Server exited with error: %+v", err)
	}
}
