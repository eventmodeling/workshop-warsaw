package router

import (
	"net/http"
	"text/template"

	app "github.com/eventmodeling/workshop-warsaw/register/app/register"
)

func getRegister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
	    panic(err)
	}

	tmpl.Execute(w, nil)
}

type postRegister struct {
	Handler app.RegisterHandler
}

func (h postRegister) Handle(w http.ResponseWriter, r *http.Request) {
	// parse cmd from request
	cmd := app.Register{}

	err := h.Handler.Execute(cmd)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("success"))
}
