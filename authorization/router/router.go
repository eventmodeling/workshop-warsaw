package router

import (
	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", login)

	return r
}
