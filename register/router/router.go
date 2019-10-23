package router

import (
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/register", getRegister)
	r.Post("/register", postRegister)

	return r
}
