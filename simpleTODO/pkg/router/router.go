package router

import (
	"fmt"
	"net/http"
	"strings"
	"todoproject/pkg/admin/adminhome"
	"todoproject/pkg/admin/adminlogin"
	"todoproject/pkg/session"
	"todoproject/pkg/user/userhome"
	"todoproject/pkg/user/userlogin"
	"todoproject/pkg/user/usermessages"
	"todoproject/pkg/user/usersignup"
	"todoproject/pkg/user/usertask"
	"todoproject/pkg/user/useruser"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if strings.HasPrefix(urlPath, "/tasks/") {
		//dedicated to tasks
		if urlPath == "/tasks/createtaskprocess" {
			usertask.CreateTaskProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/tasks/deletetask") {
			usertask.DeleteTaskProcessor(w, r)
		} else if urlPath == "/tasks/edittaskprocessor" {
			usertask.UpdateTaskProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/tasks/edittask") {
			usertask.UpdateTaskPageHandler(w, r)
		}
	} else if strings.HasPrefix(urlPath, "/users/") {
		//dedicated to users
		if urlPath == "/users/signup" {
			usersignup.SignupPageHander(w, r)
		} else if urlPath == "/users/home" {
			userhome.HomePageHandler(w, r)
		} else if urlPath == "/users/logout" {
			session.Logout(w, r)
		} else if urlPath == "/users/login" {
			userlogin.LoginPageHandler(w, r)
		} else if urlPath == "/users/signupprocess" {
			usersignup.SignupProcessHandler(w, r)
		} else if urlPath == "/users/loginprocess" {
			userlogin.LoginProcessHandler(w, r)
		} else if urlPath == "/users/edituserprocessor" {
			useruser.UpdateUserProcessor(w, r)
		} else if strings.HasPrefix(urlPath, "/users/editaccount") {
			useruser.UpdateUserPageHandler(w, r)
		} else if urlPath == "/users/messages" {
			usermessages.UserMessagePageHandler(w, r)
		} else if urlPath == "/users/createmessageprocessor" {
			usermessages.CreateUserMessageProcessor(w, r)
		}
	} else if strings.HasPrefix(urlPath, "/admin") {
		if urlPath == "/admin/login" {
			adminlogin.AdminLoginPageHandler(w, r)
		} else if urlPath == "/admin/home" {
			adminhome.AdminHomePageHandler(w, r)
		}
	} else {
		fmt.Fprintf(w, "Page Not Found")
	}
}
