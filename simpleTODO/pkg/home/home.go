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

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		username, userid := session.WhoIsThis(databasetools.DataBase, cookie.Value)
		template, _ := template.ParseFiles("../../pkg/home/template/home.html")
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "title", "description", "isDone"}, "tasks", [][]string{{"author", username}}, [][]string{})
		tasks := task.ReadTask(query, arguments)
		Data := dataToSend{Username: username, Userid: userid, Tasks: tasks}
		template.Execute(w, Data)
	}
}
