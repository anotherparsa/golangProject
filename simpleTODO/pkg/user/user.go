package user

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
)

//CRUD
//Create
func CreateUser(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("user has been created")
}

//Read
func ReadUser(query string, arguments []interface{}) []models.User {
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

//UPDATE
//page handler
func UpdateUserPageHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../../pkg/user/template/edituser.html")
	if err != nil {
		fmt.Println(err)
	}
	userId := strings.TrimPrefix(r.URL.Path, "/editaccount/")
	Query, arguments := databasetools.QuerryMaker("select", []string{"id", "username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{{"id", userId}}, [][]string{})
	user := ReadUser(Query, arguments)
	template.Execute(w, user[0])
}

//processing
func UpdateUserProcessor(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Query, arguments := databasetools.QuerryMaker("update", []string{"username", "password", "firstName", "lastName", "email", "phoneNumber"}, "users", [][]string{}, [][]string{{"username", r.Form.Get("username")}, {"firstName", r.Form.Get("firstName")}, {"lastName", r.Form.Get("lastName")}, {"email", r.Form.Get("email")}, {"phoneNumber", r.Form.Get("phoneNumber")}})
	UpdateUser(Query, arguments)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

//applying in the database{
func UpdateUser(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User Has been updated")
}
