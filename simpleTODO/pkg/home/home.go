package home

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/task"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		session_id := cookie.Value
		username := databasetools.WhoIsThis(databasetools.DB, session_id)
		template, _ := template.ParseFiles("../../pkg/home/template/home.html")
		tasks, err := task.GetUsersTask(databasetools.DB, username)
		if err != nil {
			fmt.Println(err)
		}
		template.Execute(w, tasks)

	}
}
