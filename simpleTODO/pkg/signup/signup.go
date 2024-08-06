package signup

import (
	"html/template"
	"net/http"
)

func SignupPageHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../../pkg/signup/template/signup.html")
	t.Execute(w, nil)
}
