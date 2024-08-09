package databasetools

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "testuser"
	password = "testpass"
	hostname = "localhost"
	port     = "3306"
	database = "users"
)

func connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	return sql.Open("mysql", dsn)
}

var DB *sql.DB

func Create_database() {
	DB, _ = connect()
}

//CRUD
//Create
func CreateUser(db *sql.DB, userId string, username string, password string, first_name string, last_name string, email string, phone_number string) {
	_, err := db.Exec("INSERT INTO users (userId, username, password, firstName, lastName, email, phoneNumber) VALUES (?, ?, ?, ?, ?, ?, ?)", userId, username, password, first_name, last_name, email, phone_number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Record created successfully")
}

func CreateSession(db *sql.DB, session_id string, user_id string) {
	_, err := db.Exec("INSERT INTO sessions (sessionId, userId) VALUES (?, ?)", session_id, user_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Record created successfully")
}

//Read
// Read records
/*func WhoIsThis(db *sql.DB,, factor string, value string) {
	rows, err := db.Query("SELECT username FROM users WHERE ?=?", factor, value)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var factor string
		if err := rows.Scan(&factor); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("username: %s", factor)
	}
}*/

func ReadSessions(db *sql.DB) {
	rows, err := db.Query("SELECT sessionId, userId FROM sessions")
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
		fmt.Printf("sessionId: %s, userId: %s\n", sessionId, userId)
	}
}
