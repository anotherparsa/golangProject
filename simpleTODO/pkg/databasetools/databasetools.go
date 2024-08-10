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
	fmt.Println("User created successfully")
}

func CreateSession(db *sql.DB, session_id string, user_id string) {
	_, err := db.Exec("INSERT INTO sessions (sessionId, userId) VALUES (?, ?)", session_id, user_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Session created successfully")
}

func CreateTasks(db *sql.DB, author string, priority string, title string, description string) {
	_, err := db.Exec("INSERT INTO tasks (author, priority, title, description, isdone) VALUES (?, ?, ?, ?, ?)", author, priority, title, description, "0")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Task created successfully")
}

//Read
// Read records
func WhoIsThis(db *sql.DB, session_id string) string {
	var user_id string
	var usrename string
	rows, err := db.Query("SELECT userId FROM sessions WHERE sessionId=?", session_id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Passed Session id %v \ncorresponding user id %v \n", session_id, user_id)
	}

	rows, err = db.Query("SELECT username FROM users WHERE userId=?", user_id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&usrename); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Retrieved userid id %v \ncorresponding username %v \n", user_id, username)
	}
	return username
}

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

//Delete
func DeleteTask(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Task created successfully")
}