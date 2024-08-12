package login

import (
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../../pkg/login/template/login.html")
	t.Execute(w, nil)
}
