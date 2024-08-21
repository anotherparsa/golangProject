package router

import (
	"fmt"
	"net/http"
	"strings"
	"todoproject/pkg/home"
	"todoproject/pkg/login"
	"todoproject/pkg/session"
	"todoproject/pkg/signup"
	"todoproject/pkg/task"
	"todoproject/pkg/user"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if urlPath == "/" || urlPath == "/home" {
		home.HomePageHandler(w, r)
	} else if strings.HasPrefix(urlPath, "/tasks/") {
		//dedicated to tasks
		if urlPath == "/tasks/createtaskprocess" {
			task.CreateTaskProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/tasks/deletetask") {
			task.DeleteTaskProcessor(w, r)
		} else if urlPath == "/tasks/edittaskprocessor" {
			task.UpdateTaskProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/tasks/edittask") {
			task.UpdateTaskPageHandler(w, r)
		}
	} else if strings.HasPrefix(urlPath, "/users/") {
		//dedicated to users
		if urlPath == "/users/signup" {
			signup.SignupPageHander(w, r)
		} else if urlPath == "/users/logout" {
			session.Logout(w, r)
		} else if urlPath == "/users/login" {
			login.LoginPageHandler(w, r)
		} else if urlPath == "/users/signupprocess" {
			signup.SignupProcessHandler(w, r)
		} else if urlPath == "/users/loginprocess"{
			login.LoginProcessHandler(w,r)
		} else if urlPath == "/users/edituserprocessor" {
			user.UpdateUserProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/users/editaccount") {
			user.UpdateUserPageHandler(w, r)
		}
	} else if strings.HasPrefix(urlPath, "/admin") {
		//dedicated to admin
	} else {
		fmt.Fprintf(w, "Page Not Found")
	}
}
