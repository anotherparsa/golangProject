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

type datatosend struct {
	CSRFT string
}

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil || cookie == nil {
		//generating csrf token
		csrft := tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "signupcsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		//parsing and executing the template
		datatosend := datatosend{CSRFT: csrft}
		template, _ := template.ParseFiles("../../pkg/user/usersignup/template/usersignup.html")
		template.Execute(w, datatosend)
	} else {
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func SignupProcessHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil && cookie == nil {
		generatedCSRFT, err := r.Cookie("signupcsrft")
		//checking if the csrft cookie exist or not
		if err == nil && generatedCSRFT != nil {
			//checking if the sent csrft is the same as the generated one
			r.ParseForm()
			if generatedCSRFT.Value == r.Form.Get("csrft") {
				//checking if request method is POST or not
				if r.Method == "POST" {
					//getting form input values
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					firstName := r.Form.Get("firstName")
					lastName := r.Form.Get("lastName")
					email := r.Form.Get("email")
					phoneNumber := r.Form.Get("phoneNumber")
					//checking if forms input are valid or not
					if databasetools.ValidateUserInfoFormInputs("username", username) {
						if databasetools.ValidateUserInfoFormInputs("password", password) {
							if databasetools.ValidateUserInfoFormInputs("firstName", firstName) {
								if databasetools.ValidateUserInfoFormInputs("lastName", lastName) {
									if databasetools.ValidateUserInfoFormInputs("email", email) {
										if databasetools.ValidateUserInfoFormInputs("phoneNumber", phoneNumber) {
											query, arguments := databasetools.QueryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{{"username", username}}, [][]string{})
											user := useruser.ReadUser(query, arguments)
											//making sure there is no other user with this username
											if len(user) == 0 {
												//after validating the forms input
												password = tools.HashThis(password)
												//generating user_id and session_id
												userId := tools.GenerateUUID()
												sessionId := tools.GenerateUUID()
												//setting the session_id cookie
												http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
												//creating a user record in users table
												query, arguments := databasetools.QueryMaker("insert", []string{"userId", "username", "password", "firstName", "lastName", "email", "phoneNumber", "rule", "suspended"}, "users", [][]string{}, [][]string{{"userId", userId}, {"username", username}, {"password", password}, {"firstName", firstName}, {"lastName", lastName}, {"email", email}, {"phoneNumber", phoneNumber}, {"rule", "user"}, {"suspended", "no"}})
												useruser.CreateUser(query, arguments)
												//creating a session record in sessions table
												query, arguments = databasetools.QueryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", userId}})
												session.CreateSession(query, arguments)
												//deleting csrft token cookie
												http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
												//redirecting users to their home page
												http.Redirect(w, r, "/users/home", http.StatusSeeOther)
											} else {
												http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
												http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
											}
										} else {
											http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
											http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
										}
									} else {
										http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
										http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
									}
								} else {
									http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
									http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
								}
							} else {
								http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
								http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
							}
						} else {
							http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
							http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
						}
					} else {
						http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
						http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
					}
				} else {
					http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
					http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
				http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "signupcsrft", MaxAge: -1, Path: "/"})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}
