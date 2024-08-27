package adminhome

import (
	"html/template"
	"net/http"
	"todoproject/pkg/admin/adminstatistics"
	"todoproject/pkg/session"
)

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	//check if session_id exists or not, that menas if the user is logged in or not
	if err == nil && cookie != nil {
		_, _, _, rule, _ := session.WhoIsThis(cookie.Value)
		//checking if the logged user is admin
		if rule == "admin" {
			template, _ := template.ParseFiles("../../pkg/admin/adminhome/template/adminhome.html")
			data := adminstatistics.InitializeStaticsProcess()
			template.Execute(w, data)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}
