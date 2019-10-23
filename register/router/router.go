package router

import (
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", register)

	return r
}
