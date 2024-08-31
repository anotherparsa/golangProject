package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title       string
	Description string
	Content     string
}

var Articles []Article

func main() {
	Articles = append(Articles, Article{Title: "test title1", Description: "test description1", Content: "test content1"})
	Articles = append(Articles, Article{Title: "test title2", Description: "test description2", Content: "test content2"})
	Articles = append(Articles, Article{Title: "test title3", Description: "test description3", Content: "test content3"})

	HandleRequest()
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", ShowHomePage).Methods("GET")
	myRouter.HandleFunc("/articles", ArticlesGet).Methods("GET")
	myRouter.HandleFunc("/articles", ArticlesPost).Methods("POST")
	http.ListenAndServe(":8080", myRouter)

}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

func ArticlesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Articles)
}

func ArticlesPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is post method")
}
