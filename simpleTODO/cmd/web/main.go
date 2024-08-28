package main

import (
	"fmt"
	"net/http"
	"todoproject/pkg/databasetools"
	"todoproject/pkg/router"
)

func main() {
	//calling databasetools.connect() functin which opens a sql connection and return it.
	databasetools.CreateDatabase()
	//initializing the admin user
	//temporarly disabling
	databasetools.InitializeAdminUser()
	//calling the router
	http.HandleFunc("/", router.RoutingHandler)
	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../pkg/static/"))))

	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}
