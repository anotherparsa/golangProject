package adminlogin

import (
	"html/template"
	"net/http"
)

func AdminLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("../../pkg/admin/adminlogin/template/adminlogin.html")
	template.Execute(w, nil)
}
