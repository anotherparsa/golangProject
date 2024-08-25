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
	if err != nil || cookie == nil {
		csrft = tools.GenerateUUID()
		http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
		template, _ := template.ParseFiles("../../pkg/user/userlogin/template/userlogin.html")
		template.Execute(w, nil)
	} else {
		http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func LoginProcessHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		if r.Method == "POST" {
			sent_csrf_token, err := r.Cookie("csrft")
			if err != nil || sent_csrf_token == nil {
				http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			} else {
				if sent_csrf_token.Value != csrft {
					http.Redirect(w, r, "/users/login", http.StatusSeeOther)
				} else {
					r.ParseForm()
					username := r.Form.Get("username")
					password := r.Form.Get("password")
					if tools.ValidateFormInputs("username", username) && tools.ValidateFormInputs("password", password) {
						password = tools.HashThis(password)
						if ValidateUser(username, password) {
							sessionId := tools.GenerateUUID()
							query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"username", username}, {"password", password}}, [][]string{})
							user := useruser.ReadUser(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId, Expires: time.Now().Add(time.Hour * 168), HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
							query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", user[0].UserId}})
							session.CreateSession(query, arguments)
							http.SetCookie(w, &http.Cookie{Name: "csrft", MaxAge: -1})
							http.Redirect(w, r, "/users/home", http.StatusSeeOther)
						} else {
							http.Redirect(w, r, "/users/login", http.StatusSeeOther)
						}
					} else {
						http.Redirect(w, r, "/users/login", http.StatusSeeOther)
					}
				}
			}
		} else {
			http.Redirect(w, r, "/users/login", http.StatusMethodNotAllowed)
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
