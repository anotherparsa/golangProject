package main

import (
	"fmt"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/router"
)

func main() {
	http.HandleFunc("/", router.RoutingHandler)

	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../pkg/static/"))))

	//calling databasetools.connect() functin which opens a sql connection and return it.
	databasetools.CreateDatabase()

	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}
