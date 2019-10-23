package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/eventmodeling/workshop-warsaw/register/app/feedback"
	"github.com/eventmodeling/workshop-warsaw/register/app/register"
	"github.com/go-chi/chi"
)

func NewRouter(
	registerHandler register.RegisterHandler,
	feedbackHandler feedback.FeedbackHandler,
) chi.Router {
	r := chi.NewRouter()
	r.Get("/register", renderTemplate("register"))
	r.Post("/register", postRegister{Handler: registerHandler}.Handle)

	r.Get("/feedback", renderTemplate("feedback"))
	r.Post("/feedback", postFeedback{Handler: feedbackHandler}.Handle)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")

	FileServer(r, "/static", http.Dir(filesDir))

	return r
}
