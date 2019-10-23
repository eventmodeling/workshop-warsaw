package main

import (
	"github.com/eventmodeling/workshop-warsaw/authorization/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Print("Running authorization server")

	err := http.ListenAndServe(":80", r)

	if err != nil {
		log.Fatalf("Server exited with error: %+v", err)
	}
}
