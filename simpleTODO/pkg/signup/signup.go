package signup

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/tools"
)

type datatosend struct {
	Csrf_token string
}

var csrf_token = tools.GenerateUUID()

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	d := datatosend{Csrf_token: csrf_token}
	t, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	t.Execute(w, d)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := tools.HashThis(r.Form.Get("password"))
	firstname := r.Form.Get("firstName")
	lastname := r.Form.Get("lastName")
	phonenumber := r.Form.Get("phoneNumber")
	email := r.Form.Get("email")
	sent_csrf_token := r.Form.Get("csrf-token")
	if sent_csrf_token != csrf_token {
		fmt.Fprintf(w, "Invalid CSRF token")
	} else {
		_, err := r.Cookie("session")
		if err != nil {
			userId := tools.GenerateUUID()
			sessionId := tools.GenerateUUID()
			http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId})
			databasetools.CreateUser(databasetools.DB, userId, username, password, firstname, lastname, email, phonenumber)
			databasetools.CreateSession(databasetools.DB, sessionId, userId)
		} else {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}

}
