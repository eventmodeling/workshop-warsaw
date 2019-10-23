package router

import (
	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/go-chi/chi"
	"net/http"
	"path/filepath"
	"os"
)

func NewRouter(registerHandler register.RegisterHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/register", getRegister)
	r.Post("/register", postRegister{Handler: registerHandler}.Handle)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))	
	
	return r
}
