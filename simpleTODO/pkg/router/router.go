package router

import (
	"fmt"
	"net/http"
	"todoproject/pkg/signup"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/home":
		fmt.Fprintf(w, "This is home")
	case "/signup":
		signup.SignupPageHander(w, r)
	case "/signupprocess":
		signup.SignupProcessHandler(w, r)
	}
}
