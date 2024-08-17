package databasetools

import (
	"database/sql"
	"fmt"
	"todoproject/pkg/models"

	_ "github.com/go-sql-driver/mysql"
)

//practicing
func SelectAllUsers() []models.User {
	db, _ := connect()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	users := []models.User{}

	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			fmt.Println(err)
		}

		users = append(users, user)
	}
	return users
}

func SelectUserBasedId() models.User {
	db, _ := connect()
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", 31)
	defer db.Close()
	user := models.User{}
	err := row.Scan(&user.ID, &user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	return user

}

const (
	username = "testuser"
	password = "testpass"
	hostname = "localhost"
	port     = "3306"
	database = "users"
)

func connect() (*sql.DB, error) {
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	//testuser:testpass@tcp(localhost:3306)/users
	return sql.Open("mysql", DataSourceName)
}

var DB *sql.DB

func Create_database() {
	DB, _ = connect()
}

//this way the code is vulnerable to sql injection
//fix that
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
				stringtoadd = fmt.Sprintf("%v='%v'", columnName, value)
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
					stringtoadd = fmt.Sprintf("%v='%v', ", column, value)
					counter++
				} else {
					stringtoadd = fmt.Sprintf("%v='%v' ", column, value)
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
	query := QuerryMaker("select", []string{"userId"}, "sessions", map[string]string{"sessionId": session_id}, map[string]string{})
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			fmt.Println(err)
		}
	}
	query = QuerryMaker("select", []string{"username"}, "users", map[string]string{"userId": user_id}, map[string]string{})
	rows, err = db.Query(query)
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
