package home

import (
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
	"todoproject/pkg/task"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		username := session.WhoIsThis(databasetools.DataBase, cookie.Value)
		template, _ := template.ParseFiles("../../pkg/home/template/home.html")
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "title", "description", "isDone"}, "tasks", [][]string{{"author", username}}, [][]string{})
		tasks := task.ReadTask(query, arguments)
		template.Execute(w, tasks)
	}
}
