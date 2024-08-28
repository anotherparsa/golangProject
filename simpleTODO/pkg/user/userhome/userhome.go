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
	//check if session_id exist or not, that means if the user is logged in or not
	if err == nil && cookie != nil {
		//generating csrft
		csrft := tools.GenerateUUID()
		//setting csrft cookie
		http.SetCookie(w, &http.Cookie{Name: "createtaskcsrft", Value: csrft, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode, Path: "/"})
		username, usersid, userId, _, _ := session.WhoIsThis(cookie.Value)
		//getting loged user tasks based of thier userId as the author of the tasks.
		query, arguments := databasetools.QuerryMaker("select", []string{"id", "author", "priority", "category", "title", "description", "status"}, "tasks", [][]string{{"author", userId}}, [][]string{})
		tasks := usertask.ReadTask(query, arguments)
		data := dataToSend{Username: username, Userid: usersid, Tasks: tasks, CSRFT: csrft}
		//parsing and executing the template
		template, _ := template.ParseFiles("../../pkg/user/userhome/template/userhome.html")
		template.Execute(w, data)

	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}
