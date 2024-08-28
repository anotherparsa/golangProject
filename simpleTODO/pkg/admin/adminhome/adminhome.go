package adminhome

import (
	"html/template"
	"net/http"
	"todoproject/pkg/admin/adminstatistics"
	"todoproject/pkg/models"
	"todoproject/pkg/session"
)

type datatosend struct {
	Statistics models.Static
	Username   string
}

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exists or not, that menas if the user is logged in or not
	if err == nil && cookie != nil {
		username, _, _, rule, _ := session.WhoIsThis(cookie.Value)
		//checking if the logged user is admin
		if rule == "admin" {
			data := adminstatistics.InitializeStaticsProcess()
			datatosend := datatosend{Statistics: data, Username: username}
			template, _ := template.ParseFiles("../../pkg/admin/adminhome/template/adminhome.html")
			template.Execute(w, datatosend)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	}
}
