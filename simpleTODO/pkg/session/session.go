package session

import (
	"database/sql"
	"fmt"
)

func CreateSession(db *sql.DB, query string) {
	//_, err := db.Exec("INSERT INTO sessions (sessionId, userId) VALUES (?, ?)", session_id, user_id)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}
