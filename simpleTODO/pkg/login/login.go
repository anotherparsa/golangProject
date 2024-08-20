package login

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user"
)

var csrft string

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	csrft = tools.GenerateUUID()
	http.SetCookie(w, &http.Cookie{Name: "csrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode})
	template, _ := template.ParseFiles("../../pkg/login/template/login.html")
	template.Execute(w, nil)
}

func LoginProcessHandler(w http.ResponseWriter, r *http.Request) {
	sent_csrf_token, err := r.Cookie("csrft")
	if err != nil || sent_csrf_token == nil {
		fmt.Println("csrft not found")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if sent_csrf_token.Value != csrft {
			fmt.Println("invalid csrft")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			_, err := r.Cookie("session_id")
			if err != nil {
				r.ParseForm()
				username := r.Form.Get("username")
				password := tools.HashThis(r.Form.Get("password"))
				if ValidateUser(username, password) {
					sessionId := tools.GenerateUUID()
					query, arguments := databasetools.QuerryMaker("select", []string{"id", "userId", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"username", username}, {"password", password}}, [][]string{})
					user := user.ReadUser(query, arguments)
					http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId})
					query, arguments = databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", [][]string{}, [][]string{{"sessionId", sessionId}, {"userId", user[0].UserId}})
					session.CreateSession(query, arguments)
					http.Redirect(w, r, "/home", http.StatusSeeOther)
				} else {
					fmt.Println("User not found in login process handler ")
				}
			}
		}
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
