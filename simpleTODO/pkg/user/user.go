package user

import (
	"database/sql"
	"fmt"
	"todoproject/pkg/models"
)

//CRUD
//Create
func CreateUser(db *sql.DB, userId string, username string, password string, firstName string, lastName string, email string, phoneNumber string) {
	_, err := db.Exec("INSERT INTO users (userId, username, password, firstName, lastName, email, phoneNumber) VALUES (?, ?, ?, ?, ?, ?, ?)", userId, username, password, firstName, lastName, email, phoneNumber)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadUser(db *sql.DB, factor string, value string) models.User {
	user := models.User{}
	rows, err := db.Query("SELECT userId, username, password, firstName, lastName, email, phoneNumber WHERE ?=?", factor, value)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			fmt.Println(err)
		}
	}
	return user
}
