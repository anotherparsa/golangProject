package session

import (
	"database/sql"
	"fmt"
	"log"
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
	fmt.Println("session has been created")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", MaxAge: -1, Path: "/"})
	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}

//finding a corespounding user_id user's id and username from a session id.
func WhoIsThis(session_id string) (string, string, string) {
	var user_id string
	var username string
	var users_id string
	query, arguments := databasetools.QuerryMaker("select", []string{"userId"}, "sessions", [][]string{{"sessionId", session_id}}, [][]string{})
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			fmt.Println(err)
		}
	}
	query, arguments = databasetools.QuerryMaker("select", []string{"username", "id"}, "users", [][]string{{"userId", user_id}}, [][]string{})
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
		if err := rows.Scan(&username, &users_id); err != nil {
			fmt.Println(err)
		}
	}
	return username, users_id, user_id
}

func ReadSessions(database *sql.DB) {
	rows, err := database.Query("SELECT sessionId, userId FROM sessions")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var sessionId string
		var userId string
		if err := rows.Scan(&sessionId, &userId); err != nil {
			fmt.Println(err)
		}
	}
}

func ReturnUsersUserID(sessionId string) string {
	var user_id string
	query, arguments := databasetools.QuerryMaker("select", []string{"userId"}, "sessions", [][]string{{"sessionId", sessionId}}, [][]string{})
	safequeyr, err := databasetools.DataBase.Prepare(query)

	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequeyr.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			fmt.Println(err)
		}
	}
	return user_id
}
