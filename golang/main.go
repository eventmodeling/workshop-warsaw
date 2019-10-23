package main

import (
	"log"
	"net/http"

	"github.com/eventmodeling/workshop-warsaw/register/app/feedback"
	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/eventmodeling/workshop-warsaw/register/infrastructure/events"
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

	feedbackHandler := feedback.FeedbackHandler{}

	r := router.NewRouter(registerHandler, feedbackHandler)
	log.Print("Running golang server")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalf("Server exited with error: %+v", err)
	}
}
