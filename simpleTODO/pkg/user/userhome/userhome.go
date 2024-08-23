package userhome

import (
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
	"todoproject/pkg/user/usertask"
)

type dataToSend struct {
	Username string
	Userid   string
	Tasks    []models.Task
}

//home page handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//check if user is logged in or not
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
	} else {
		username, usersid, userId := session.WhoIsThis(cookie.Value)
		template, _ := template.ParseFiles("../../pkg/user/userhome/template/userhome.html")
		//making query to get tasks
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "finished"}, "tasks", [][]string{{"author", userId}}, [][]string{})
		//getting tasks
		tasks := usertask.ReadTask(query, arguments)
		Data := dataToSend{Username: username, Userid: usersid, Tasks: tasks}
		template.Execute(w, Data)
	}
}
