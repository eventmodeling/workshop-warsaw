package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/eventmodeling/workshop-warsaw/register/infrastructure"
	"github.com/eventmodeling/workshop-warsaw/register/router"
)

const eventsDirectory = "/events"
const userRegisteredEventFile = "UserRegistered.json"

func main() {
	filePath := path.Join(eventsDirectory, userRegisteredEventFile)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	publisher := infrastructure.NewEventPublisher(file)

	registerHandler := register.RegisterHandler{
		Publisher: publisher,
	}

	r := router.NewRouter(registerHandler)
	log.Print("Running registration server")

	err = http.ListenAndServe(":80", r)

	if err != nil {
		log.Fatalf("Server exited with error: %+v", err)
	}
}
