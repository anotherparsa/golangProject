package router

import (
	"fmt"
	"net/http"
	"strings"
	"todoproject/pkg/home"
	"todoproject/pkg/signup"
	"todoproject/pkg/task"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if urlPath == "/" || urlPath == "/home" {
		home.HomePageHandler(w, r)
	} else if urlPath == "/signup" {
		signup.SignupPageHander(w, r)
	} else if urlPath == "/signupprocess" {
		signup.SignupProcessHandler(w, r)
	} else if urlPath == "/createtaskprocess" {
		task.CreateTaskProcessor(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/deletetask/") {
		task.DeleteTask(w, r)
	} else if urlPath == "/edittaskprocessor" {
		task.EditTask(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/edittask/") {
		task.EditTaskPageHandler(w, r)
	} else {
		fmt.Fprintf(w, "Page not found")
	}

}
