package router

import (
	app "github.com/eventmodeling/workshop-warsaw/register/app/register"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	// parse cmd from request
	cmd := app.Register{}
	h := app.RegisterHandler{}

	err := h.Execute(cmd)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("success"))
}