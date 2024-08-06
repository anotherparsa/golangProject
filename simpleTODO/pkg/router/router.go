package router

import (
	"fmt"
	"net/http"
)

func RoutingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/home":
		fmt.Fprintf(w, "This is home")
	case "/signup":
		fmt.Fprintf(w, "This is signup")
	}
}
