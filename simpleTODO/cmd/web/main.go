package main

import (
	"fmt"
	"net/http"

	// you shouldn't prefix your packages by `/pkg/`. It's an old idiom that is not used anymore.
	// every package is public by default. If you want keep it for your app only, put them under an `internal` folder.
	"todoproject/pkg/databasetools"
	"todoproject/pkg/router"
)

func main() {
	//calling databasetools.connect() functin which opens a sql connection and return it.
	// it should return a database (either the *sql.DB, or some more app oriented store object) that you can inject
	// in your other function/objects
	databasetools.CreateDatabase()
	//initializing the admin user
	//temporarly disabling
	databasetools.InitializeAdminUser()
	//calling the router
	http.HandleFunc("/", router.RoutingHandler)
	//serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../pkg/static/"))))

	fmt.Println("Running server on port 8080")
	// you could catch the error from this call. It should very rarely happen, but still,
	// at least to know what happened for your app to just stop
	http.ListenAndServe(":8080", nil)
}
