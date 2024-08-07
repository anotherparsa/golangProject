package signup

import (
	"html/template"
	"net/http"
	"todoproject/pkg/tools"
)

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	t.Execute(w, nil)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	firstname := r.Form.Get("firstName")
	lastname := r.Form.Get("lastName")
	phonenumber := r.Form.Get("phoneNumber")
	email := r.Form.Get("email")
	tools.CreateUser(tools.DB, username, password, firstname, lastname, email, phonenumber)

}
