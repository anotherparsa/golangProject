package main

import (
	"fmt"
	"net/http"
	"todoproject/pkg/router"
)

func main() {
	http.HandleFunc("/", router.RoutingHandler)

	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}
