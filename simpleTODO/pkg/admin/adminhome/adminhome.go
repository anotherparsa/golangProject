package adminhome

import (
	"html/template"
	"net/http"
	"todoproject/pkg/admin/adminstatistics"
)

func AdminHomePageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("../../pkg/admin/adminhome/template/adminhome.html")
	data := adminstatistics.InitializeStaticsProcess()
	template.Execute(w, data)
}
