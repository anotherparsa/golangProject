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
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil || cookie == nil {
		//generating csrf token
		csrft = tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
		//parsing and executing the template
		template, _ := template.ParseFiles("../../pkg/user/usersignup/template/usersignup.html")
		template.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil && cookie == nil {
		sent_csrf_token, err := r.Cookie("csrft")
		//checking if the csrft cookie exist or not
		if err == nil && sent_csrf_token != nil {
			//checking if the sent csrft is the same as the generated one
			if sent_csrf_token.Value == csrft {
				//checking if request method is POST or not
				if r.Method == "POST" {
					//getting form input values
					r.ParseForm()
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					firstName := r.Form.Get("firstName")
					lastName := r.Form.Get("lastName")
					email := r.Form.Get("email")
					phoneNumber := r.Form.Get("phoneNumber")
					//checking if forms input are valid or not
					if tools.ValidateFormInputs("username", username) {
						if tools.ValidateFormInputs("password", password) {
							if tools.ValidateFormInputs("firstName", firstName) {
								if tools.ValidateFormInputs("lastName", lastName) {
									if tools.ValidateFormInputs("email", email) {
										if tools.ValidateFormInputs("phoneNumber", phoneNumber) {
											//after validating the forms input
											password = tools.HashThis(password)
											//generating user_id and session_id
											userId := tools.GenerateUUID()
											sessionId := tools.GenerateUUID()
											//setting the session_id cookie
											http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
											//creating a user record in users table
											query, arguments := databasetools.QuerryMaker("insert", []string{"userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{}, [][]string{{"userId", userId}, {"username", username}, {"password", tools.HashThis(password)}, {"firstName", firstName}, {"lastName", lastName}, {"email", email}, {"phoneNumber", phoneNumber}, {"rule", "user"}, {"suspended", "no"}})
											useruser.CreateUser(query, arguments)
											//creating a session record in sessions table
											query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", userId}})
											session.CreateSession(query, arguments)
											//deleting csrft token cookie
											http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
											//redirecting users to their home page
											http.Redirect(w, r, "/users/home", http.StatusSeeOther)
										} else {
											http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
											http.Redirect(w, r, "/users/login", http.StatusSeeOther)
										}
									} else {
										http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
										http.Redirect(w, r, "/users/login", http.StatusSeeOther)
									}
								} else {
									http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
									http.Redirect(w, r, "/users/login", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
								http.Redirect(w, r, "/users/login", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
							http.Redirect(w, r, "/users/login", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
						http.Redirect(w, r, "/users/login", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
					http.Redirect(w, r, "/users/login", http.StatusMethodNotAllowed)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
				http.Redirect(w, r, "/users/signup", http.StatusUnauthorized)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
			http.Redirect(w, r, "/users/signup", http.StatusUnauthorized)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}
