package usersignup

import (
	"html/template"
	"net/http"
	"time"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user/useruser"
)

var csrft string

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		csrft = tools.GenerateUUID()
		http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
		template, _ := template.ParseFiles("../../pkg/user/usersignup/template/usersignup.html")
		template.Execute(w, nil)
	} else {
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		if r.Method == "POST" {
			sent_csrf_token, err := r.Cookie("csrft")
			if err != nil || sent_csrf_token == nil {
				http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
			} else {
				if sent_csrf_token.Value != csrft {
					http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
				} else {
					r.ParseForm()
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					firstName := r.Form.Get("firstName")
					lastName := r.Form.Get("lastName")
					email := r.Form.Get("email")
					phoneNumber := r.Form.Get("phoneNumber")
					if tools.ValidateFormInputs("username", username) && tools.ValidateFormInputs("password", password) && tools.ValidateFormInputs("firstname", firstName) && tools.ValidateFormInputs("lastname", lastName) && tools.ValidateFormInputs("email", email) && tools.ValidateFormInputs("phonenumber", phoneNumber) {
						userId := tools.GenerateUUID()
						sessionId := tools.GenerateUUID()
						http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
						query, arguments := databasetools.QuerryMaker("insert", []string{"userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{}, [][]string{{"userId", userId}, {"username", username}, {"password", tools.HashThis(password)}, {"firstName", firstName}, {"lastName", lastName}, {"email", email}, {"phoneNumber", phoneNumber}, {"rule", "user"}, {"suspended", "no"}})
						useruser.CreateUser(query, arguments)
						query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", userId}})
						session.CreateSession(query, arguments)
						http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
					} else {
						http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
					}
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
				}
			}
		} else {
			http.Redirect(w, r, "/users/signup", http.StatusMethodNotAllowed)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}
