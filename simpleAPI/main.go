package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc`
	Content string `json:"content"`
}

type Articles []Article

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
