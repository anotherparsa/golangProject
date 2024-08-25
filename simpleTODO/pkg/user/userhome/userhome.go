package userhome

import (
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
	"todoproject/pkg/tools"
	"todoproject/pkg/user/usertask"
)

type dataToSend struct {
	Username string
	Userid   string
	CSRFT    string
	Tasks    []models.Task
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/users/signup", http.StatusSeeOther)
	} else {
		csrft := tools.GenerateUUID()
		http.SetCookie(w, &http.Cookie{Name: "homecsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		template, _ := template.ParseFiles("../../pkg/user/userhome/template/userhome.html")
		username, usersid, userId := session.WhoIsThis(cookie.Value)
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "finished"}, "tasks", [][]string{{"author", userId}}, [][]string{})
		tasks := usertask.ReadTask(query, arguments)
		data := dataToSend{Username: username, Userid: usersid, Tasks: tasks, CSRFT: csrft}
		template.Execute(w, data)
	}
}
