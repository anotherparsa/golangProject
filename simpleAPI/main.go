package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	http.HandleFunc("/", ShowHomePage)
	http.HandleFunc("/articles", ShowArticles)
	http.ListenAndServe(":8080", nil)

}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

func ShowArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}
