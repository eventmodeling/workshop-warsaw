package router

import (
	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/go-chi/chi"
)

func NewRouter(registerHandler register.RegisterHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/register", getRegister)
	r.Post("/register", postRegister{Handler: registerHandler}.Handle)

	return r
}
