package user

import (
	"database/sql"
	"fmt"
)

//CRUD
//Create
func CreateUser(db *sql.DB, userId string, username string, password string, firstName string, lastName string, email string, phoneNumber string) {
	_, err := db.Exec("INSERT INTO users (userId, username, password, firstName, lastName, email, phoneNumber) VALUES (?, ?, ?, ?, ?, ?, ?)", userId, username, password, firstName, lastName, email, phoneNumber)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User Created")
	}
}
