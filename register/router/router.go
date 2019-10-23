package router

import (
	"github.com/go-chi/chi"
	"net/http"
	"path/filepath"
	"os"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/register", getRegister)
	r.Post("/register", postRegister)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))	
	
	return r
}
