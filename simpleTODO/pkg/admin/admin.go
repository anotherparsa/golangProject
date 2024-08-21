package admin

import (
	"html/template"
	"net/http"
	"todoproject/pkg/admin/statics"
)

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("../../pkg/admin/template/adminhome.html")
	data := statics.InitializeStaticsProcess()
	template.Execute(w, data)
}
