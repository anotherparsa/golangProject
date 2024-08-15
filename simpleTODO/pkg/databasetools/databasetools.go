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

func ReadQuerryMaker(coulumns []string, table string, conditions map[string]string) string {
	var stringtoadd string
	Querry := "SELECT "
	//creating columns part
	for i := 0; i < len(coulumns); i++ {
		if i <= len(coulumns)-2 {
			stringtoadd = fmt.Sprintf("%v, ", coulumns[i])
		} else {
			stringtoadd = fmt.Sprintf("%v ", coulumns[i])
		}
		Querry += stringtoadd
	}
	Querry += fmt.Sprintf("FROM %v", table)
	fmt.Println(Querry)
	//condition part
	if len(conditions) != 0 {
		stringtoadd = ""
		Querry += " WHERE "
		for columnName, value := range conditions {
			stringtoadd = fmt.Sprintf("%v=%v", columnName, value)
		}
	}
	Querry += stringtoadd
	fmt.Println(Querry)
	return Querry

}

func CreateSession(db *sql.DB, session_id string, user_id string) {
	_, err := db.Exec("INSERT INTO sessions (sessionId, userId) VALUES (?, ?)", session_id, user_id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Session created successfully")
}

//Read
// Read records
func WhoIsThis(db *sql.DB, session_id string) string {
	fmt.Printf("Passed session id : %v\n", session_id)
	var user_id string
	var username string
	rows, err := db.Query("SELECT userId FROM sessions WHERE sessionId=?", session_id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Found user id : %v\n", user_id)
	rows, err = db.Query("SELECT username FROM users WHERE userId=?", user_id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&username); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Found user name : %v\n", username)
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
	}
}

//Delete
func DeleteTask(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Task Deleted successfully")
}

//update task
func EditTask(db *sql.DB, id string, newTitle string, newDescription string, newPriority string) {
	_, err := db.Exec("UPDATE tasks set priority=?, title=?, description=? WHERE id=?", newPriority, newTitle, newDescription, id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("we've go an error")
	}
	fmt.Println("Task Updated successfully")
}

func ValidateUser(db *sql.DB, username string, password string) bool {
	rows, err := db.Query("SELECT password FROM users where username=?", username)
	fmt.Printf("The user name passed in validate user is %v \n and the password is %v \n", username, password)
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
	fmt.Printf("the passed password is %v \nand the stored password is %v\n", password, storedPassword)
	return storedPassword == password

}

func GetUsersUserid(db *sql.DB, username string) string {
	var user_id string
	rows, err := db.Query("SELECT userId FROM users WHERE username=?", username)

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
