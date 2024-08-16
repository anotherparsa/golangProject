package login

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../../pkg/login/template/login.html")
	t.Execute(w, nil)
}

func LoginProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := tools.HashThis(r.Form.Get("password"))
	if ValidateUser(databasetools.DB, username, password) {
		sessionId := tools.GenerateUUID()
		userId := databasetools.GetUsersUserid(databasetools.DB, username)
		http.SetCookie(w, &http.Cookie{Name: "session_id", Value: sessionId})
		query := databasetools.QuerryMaker("insert", []string{"sessionId", "userId"}, "sessions", map[string]string{}, map[string]string{"sessionId": sessionId, "userId": userId})
		session.CreateSession(databasetools.DB, query)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		fmt.Println("User not found in login process handler ")
	}

}
func ValidateUser(db *sql.DB, username string, password string) bool {
	rows, err := db.Query("SELECT password FROM users where username=?", username)
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
