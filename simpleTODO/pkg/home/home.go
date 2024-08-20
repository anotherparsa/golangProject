package home

import (
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
	"todoproject/pkg/task"
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
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		username, usersid, userId := session.WhoIsThis(cookie.Value)
		template, _ := template.ParseFiles("../../pkg/home/template/home.html")
		//making query to get tasks
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "title", "description", "isDone"}, "tasks", [][]string{{"author", userId}}, [][]string{})
		//getting tasks
		tasks := task.ReadTask(query, arguments)
		Data := dataToSend{Username: username, Userid: usersid, Tasks: tasks}
		template.Execute(w, Data)
	}
}
