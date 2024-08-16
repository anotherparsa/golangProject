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

func QuerryMaker(operation string, coulumns []string, table string, conditions map[string]string, values map[string]string) string {
	//read part
	if operation == "select" {
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
		//condition part
		if len(conditions) != 0 {
			stringtoadd = ""
			Querry += " WHERE "
			for columnName, value := range conditions {
				stringtoadd = fmt.Sprintf("%v=%v", columnName, value)
			}
		}
		Querry += stringtoadd
		return Querry
		//update part
	} else if operation == "update" {
		var stringtoadd string
		Query := "UPDATE " + table + " SET "
		if len(values) != 0 {
			stringtoadd = ""
			counter := 0
			for column, value := range values {
				if counter <= len(values)-2 {
					stringtoadd = fmt.Sprintf("%v=%v, ", column, value)
					counter++
				} else {
					stringtoadd = fmt.Sprintf("%v=%v ", column, value)
				}
				Query += stringtoadd
			}
		}
		if len(conditions) != 0 {
			stringtoadd = ""
			Query += "WHERE "
			for columnName, value := range conditions {
				stringtoadd = fmt.Sprintf("%v=%v", columnName, value)
				Query += stringtoadd
			}
		}
		return Query
	} else if operation == "insert" {
		var stringtoadd string
		Query := fmt.Sprintf("INSERT INTO %v ( ", table)
		for i := 0; i < len(coulumns); i++ {
			if i <= len(coulumns)-2 {
				stringtoadd = fmt.Sprintf("%v, ", coulumns[i])
			} else {
				stringtoadd = fmt.Sprintf("%v ", coulumns[i])
			}
			Query += stringtoadd
		}
		Query += ") VALUES ( "
		if len(values) != 0 {
			stringtoadd = ""
			counter := 0
			for _, value := range values {
				if counter <= len(values)-2 {
					stringtoadd = fmt.Sprintf("'%v', ", value)
					counter++
				} else {
					stringtoadd = fmt.Sprintf("'%v' ", value)
				}
				Query += stringtoadd
			}
		}
		Query += ")"
		return Query
	}
	return "invalid operation"
}

//Read
// Read records
func WhoIsThis(db *sql.DB, session_id string) string {
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
}

//update task
func EditTask(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		fmt.Println("we've go an error")
	}
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
