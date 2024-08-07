package signup

import (
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/tools"
)

type datatosend struct {
	Csrf_token string
}

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	d := datatosend{Csrf_token: tools.GenerateUUID()}
	t, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	t.Execute(w, d)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	firstname := r.Form.Get("firstName")
	lastname := r.Form.Get("lastName")
	phonenumber := r.Form.Get("phoneNumber")
	email := r.Form.Get("email")
	databasetools.CreateUser(databasetools.DB, username, password, firstname, lastname, email, phonenumber)

}
