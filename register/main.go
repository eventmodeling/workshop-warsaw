package main

import (
	"log"
	"net/http"

	"github.com/eventmodeling/workshop-warsaw/register/router"
)

func main() {
	r := router.NewRouter()
	log.Print("Running registration server")

	err := http.ListenAndServe(":80", r)

	if err != nil {
		log.Fatalf("Server exited with error: %+v", err)
	}
}
