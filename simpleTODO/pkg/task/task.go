package task

import (
	"net/http"
	"todoproject/pkg/databasetools"
)

func CreateTaskProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else {
		r.ParseForm()
		author := databasetools.WhoIsThis(databasetools.DB, cookie.Value)
		priority := r.Form.Get("priority")
		title := r.Form.Get("title")
		description := r.Form.Get("description")
		databasetools.CreateTasks(databasetools.DB, author, priority, title, description)

	}
}
