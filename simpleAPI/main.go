package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	HandleRequests()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
