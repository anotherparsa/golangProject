package signup

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/tools"
	"todoproject/pkg/user"
)

var csrft string

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	csrft = tools.GenerateUUID()
	http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
	template, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	template.Execute(w, nil)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := tools.HashThis(r.Form.Get("password"))
	firstname := r.Form.Get("firstName")
	lastname := r.Form.Get("lastName")
	phonenumber := r.Form.Get("phoneNumber")
	email := r.Form.Get("email")
	sent_csrf_token, err := r.Cookie("csrft")
	if err != nil || sent_csrf_token == nil {
		fmt.Println(err)
	}
	if sent_csrf_token.Value != csrft {
		fmt.Fprintf(w, "Invalid CSRF token")
	} else {
		_, err := r.Cookie("session")
		if err != nil {
			userId := tools.GenerateUUID()
			sessionId := tools.GenerateUUID()
			http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId})
			user.CreateUser(databasetools.DB, userId, username, password, firstname, lastname, email, phonenumber)
			databasetools.CreateSession(databasetools.DB, sessionId, userId)
			http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		} else {
			http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

}
