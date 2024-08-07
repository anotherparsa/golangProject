package main

import (
	"fmt"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/router"
)

func main() {
	http.HandleFunc("/", router.RoutingHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../pkg/static/"))))

	databasetools.Create_database()
	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}
