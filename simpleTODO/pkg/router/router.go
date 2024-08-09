package router

import (
	"net/http"
	"todoproject/pkg/home"
	"todoproject/pkg/signup"
	"todoproject/pkg/task"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/home":
		home.HomePageHandler(w, r)
	case "/signup":
		signup.SignupPageHander(w, r)
	case "/signupprocess":
		signup.SignupProcessHandler(w, r)
	case "/createtaskprocess":
		task.CreateTaskProcessor(w, r)
	}
}
