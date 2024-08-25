package userlogin

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user/useruser"
)

var csrft string

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil || cookie == nil {
		//generating csrf token
		csrft = tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
		//parsing and executing the template
		template, _ := template.ParseFiles("../../pkg/user/userlogin/template/userlogin.html")
		template.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func LoginProcessHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exist or not, that means if the user is logged in or not
	if err != nil && cookie == nil {
		sent_csrf_token, err := r.Cookie("csrft")
		//checking if the csrft cookie exist or not
		if err == nil && sent_csrf_token != nil {
			//checking if the sent csrft is the same as the generated one
			if sent_csrf_token.Value == csrft {
				//checking if the request method is POST or not
				if r.Method == "POST" {
					//getting form input values
					r.ParseForm()
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					//checking if form inputs are valid or not
					if tools.ValidateFormInputs("username", username) && tools.ValidateFormInputs("password", password) {
						//after form inputs have been validated.
						password = tools.HashThis(password)
						//validating the user in the database
						if ValidateUser(username, password) {
							//generating a session_id
							sessionId := tools.GenerateUUID()
							//getting user to set a session_id corresponding to their userId
							query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"username", username}, {"password", password}}, [][]string{})
							user := useruser.ReadUser(query, arguments)
							//setting the session_id cookie
							http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
							//creating a session record in the session table
							query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", user[0].UserId}})
							session.CreateSession(query, arguments)
							//deleting csrft token cookie
							http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
							//redirecting user to their home page
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
					http.Redirect(w, r, "/users/login", http.StatusMethodNotAllowed)
				}
			} else {
				http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
				http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
			http.Redirect(w, r, "/users/login", http.StatusUnauthorized)
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}
func ValidateUser(username string, password string) bool {
	rows, err := databasetools.DataBase.Query("SELECT password FROM users where username=?", username)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	if !rows.Next() {
		fmt.Println("User Not found in validate user")
	}
	var storedPassword string
	err = rows.Scan(&storedPassword)
	if err != nil {
		fmt.Println(err)
	}
	return storedPassword == password

}
