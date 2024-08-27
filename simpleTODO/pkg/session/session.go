package session

import (
	"fmt"
	"net/http"
	"todoproject/pkg/databasetools"
)

func CreateSession(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)

	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)

	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "homecsrft", MaxAge: -1, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "session_id", MaxAge: -1, Path: "/"})
	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}

func WhoIsThis(session_id string) (string, string, string, string, string) {
	var username string
	var users_id string
	var user_id string
	var suspended string
	var rule string

	//getting the userId of the logged user corresponding to the session_id
	query, arguments := databasetools.QuerryMaker("select", []string{"userId"}, "sessions", [][]string{{"sessionId", session_id}}, [][]string{})
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&user_id)
		if err != nil {
			fmt.Println(err)
		}
	}
	//getting user's id, user's username', user's rule and user's suspend status
	query, arguments = databasetools.QuerryMaker("select", []string{"id", "username", "rule", "suspended"}, "users", [][]string{{"userId", user_id}}, [][]string{})
	safequery, err = databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err = safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&users_id, &username, &rule, &suspended); err != nil {
			fmt.Println(err)
		}
	}
	return username, users_id, user_id, rule, suspended
}
