package router

import (
	app "github.com/eventmodeling/workshop-warsaw/authorization/app/login"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	// parse cmd from request
	cmd := app.Login{}
	h := app.LoginHandler{}

	err := h.Execute(cmd)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("success"))
}