package router

import (
	"errors"
	"net/http"
	"text/template"

	app "github.com/eventmodeling/workshop-warsaw/register/app/register"
)

type registerData struct {
	name     string
	password string
	email    string
}

func getRegister(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

type postRegister struct {
	Handler app.RegisterHandler
}

func (h postRegister) Handle(w http.ResponseWriter, r *http.Request) {
	data, err := parseRegisterData(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// parse cmd from request
	cmd := app.Register{
		Name:     data.name,
		Email:    data.email,
		Password: data.password,
	}
	err = h.Handler.Execute(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("success"))
}

func parseRegisterData(r *http.Request) (registerData, error) {
	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")

	if name == "" {
		return registerData{}, errors.New("name is empty")
	}
	if password == "" {
		return registerData{}, errors.New("password is empty")
	}
	if email == "" {
		return registerData{}, errors.New("email is empty")
	}

	return registerData{
		name:     name,
		password: password,
		email:    email,
	}, nil
}
