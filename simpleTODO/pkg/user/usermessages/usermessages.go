package usermessages

import (
	"fmt"
	"html/template"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/session"
)

func UserMessagePageHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		template, _ := template.ParseFiles("../../pkg/user/usermessages/template/usermessage.html")
		template.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/users/home", http.StatusSeeOther)
	}
}

func CreateUserMessageProcessor(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie == nil {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	} else {
		if r.Method == "POST" {
			r.ParseForm()
			author, _, _ := session.WhoIsThis(cookie.Value)
			query, arguments := databasetools.QuerryMaker("insert", []string{"author", "priority", "category", "title", "description", "finished"}, "messages", [][]string{}, [][]string{{"author", author}, {"priority", r.Form.Get("priority")}, {"category", r.Form.Get("category")}, {"title", r.Form.Get("title")}, {"description", r.Form.Get("description")}, {"finished", "unfinished"}})
			CreateMessage(query, arguments)
			http.Redirect(w, r, "/users/home", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/users/home", http.StatusMethodNotAllowed)
		}
	}

}

func CreateMessage(query string, arguments []interface{}) {
	safequery, err := databasetools.DataBase.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	_, err = safequery.Exec(arguments...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message has been created")
}
