package router

import (
	app "github.com/eventmodeling/workshop-warsaw/register/app/register"
	"text/template"
	"net/http"
)

func getRegister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
	    panic(err)
	}

	tmpl.Execute(w, nil)
}

func postRegister(w http.ResponseWriter, r *http.Request) {
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
