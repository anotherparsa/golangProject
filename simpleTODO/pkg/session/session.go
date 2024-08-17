package session

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	databasetool "todoproject/pkg/databasetools"
)

func CreateSession(db *sql.DB, query string) {
	//_, err := db.Exec("INSERT INTO sessions (sessionId, userId) VALUES (?, ?)", session_id, user_id)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_id", MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func WhoIsThis(database *sql.DB, session_id string) string {
	var user_id string
	var username string
	query, arguments := databasetool.QuerryMaker("select", []string{"userId"}, "sessions", map[string]string{"sessionId": session_id}, map[string]string{})
	safequery, err := database.Prepare(query)
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
	query, arguments = databasetool.QuerryMaker("select", []string{"username"}, "users", map[string]string{"userId": user_id}, map[string]string{})
	safequery, err = database.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err = safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&username); err != nil {
			fmt.Println(err)
		}
	}
	return username
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
