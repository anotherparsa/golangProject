package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		{
			Title:   "test title",
			Desc:    "test desc",
			Content: "test content",
		},
	}
	fmt.Fprintf(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page Endpoint hit")
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/articles", AllArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	HandleRequest()
}
