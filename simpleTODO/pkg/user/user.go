package user

import (
	"database/sql"
	"fmt"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

//CRUD
//Create
func CreateUser(database *sql.DB, query string, arguments []interface{}) {
	fmt.Println("error in 1")
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println("error in 2")
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err == nil {
		fmt.Println("error in 3")
		fmt.Println(err)
	}
	fmt.Println("error in 4")
}

//Read
func ReadUser(database *sql.DB, query string, arguments []interface{}) []models.User {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := safequery.Query(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	user := models.User{}
	users := []models.User{}

	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	return users
}
