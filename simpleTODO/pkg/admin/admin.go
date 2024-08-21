package admin

import (
	"html/template"
	"net/http"
)

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("../../pkg/admin/template/adminhome.html")
	template.Execute(w, nil)
}
