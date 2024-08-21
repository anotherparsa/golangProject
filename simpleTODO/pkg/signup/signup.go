package signup

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user"
)

var csrft string

//generating csrft and set it as a cookie to be retrieved in login process handler and be checked.
func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	csrft = tools.GenerateUUID()
	http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
	template, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	template.Execute(w, nil)
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sent_csrf_token, err := r.Cookie("csrft")
		if err != nil || sent_csrf_token == nil {
			fmt.Println("csrft not found")
			http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
		} else {
			if sent_csrf_token.Value != csrft {
				fmt.Println("invalid csrft")
				http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
			} else {
				_, err := r.Cookie("session_id")
				if err != nil {
					r.ParseForm()
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					firstName := r.Form.Get("firstName")
					lastName := r.Form.Get("lastName")
					email := r.Form.Get("email")
					phoneNumber := r.Form.Get("phoneNumber")

					if tools.ValidateSignupFormInputs(username, password, firstName, lastName, email, phoneNumber) {
						userId := tools.GenerateUUID()
						sessionId := tools.GenerateUUID()
						http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})

						//Create user query
						query, arguments := databasetools.QuerryMaker("insert", []string{"userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{}, [][]string{{"userId", userId}, {"username", username}, {"password", tools.HashThis(password)}, {"firstName", firstName}, {"lastName", lastName}, {"email", email}, {"phoneNumber", phoneNumber}})
						user.CreateUser(query, arguments)

						//Create Session query
						query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", userId}})
						session.CreateSession(query, arguments)

						http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
					} else {
						fmt.Println("Invalid inputs")
						http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
					http.Redirect(w, r, "/users/home", http.StatusSeeOther)
				}
				http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
				http.Redirect(w, r, "/users/home", http.StatusSeeOther)
			}
		}
	} else {
		fmt.Println("wrong method")
		http.Redirect(w, r, "/users/signup", http.StatusMethodNotAllowed)
	}
}
