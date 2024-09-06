package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Claims struct {
	Username string `json:"username"`
	UsersId  string `json:"usersid`
	UserId   string `json:"userid"`
	Expires  int64  `json:"expire"`
}

type Article struct {
	ID          string
	Title       string
	Description string
	Content     string
}

var Articles []Article

func main() {
	Articles = append(Articles, Article{ID: "1", Title: "test title1", Description: "test description1", Content: "test content1"})
	Articles = append(Articles, Article{ID: "2", Title: "test title2", Description: "test description2", Content: "test content2"})
	Articles = append(Articles, Article{ID: "3", Title: "test title3", Description: "test description3", Content: "test content3"})
	Articles = append(Articles, Article{ID: "4", Title: "test title4", Description: "test description4", Content: "test content4"})

	HandleRequest()
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", ShowHomePage).Methods("GET")
	myRouter.HandleFunc("/articles", ShowArticles).Methods("GET")
	myRouter.HandleFunc("/articles", AddNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", ShowArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/encode", Encode).Methods("GET")

	http.ListenAndServe(":8080", myRouter)

}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

func ShowArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Articles)
}

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	inputs := mux.Vars(r)
	articleID := inputs["id"]
	for _, article := range Articles {
		if article.ID == articleID {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(article)
		}
	}
}

func AddNewArticle(w http.ResponseWriter, r *http.Request) {
	RequestBody, _ := io.ReadAll(r.Body)
	article := Article{}
	_ = json.Unmarshal(RequestBody, &article)
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Article has been appended")

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	inputs := mux.Vars(r)
	articleId := inputs["id"]
	for index, article := range Articles {
		if article.ID == articleId {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Article has been deleted")
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	inputs := mux.Vars(r)
	articleId := inputs["id"]
	RequestBody, _ := io.ReadAll(r.Body)
	SingleArticle := Article{}
	_ = json.Unmarshal(RequestBody, &SingleArticle)
	for index, article := range Articles {
		if article.ID == articleId {
			Articles[index].Title = SingleArticle.Title
			Articles[index].Description = SingleArticle.Description
			Articles[index].Content = SingleArticle.Content
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Article updated")
		}
	}
}

func Encode(w http.ResponseWriter, r *http.Request) {
	//marshal == struct or any other data type to json encoded byte slice
	//unmarshal == json encoded byte slice to struct or any other data type
	A1 := Article{ID: "6", Title: "TestTitle"}
	A1ByteArray, _ := json.Marshal(A1)
	fmt.Println(A1ByteArray)
	fmt.Println(string(A1ByteArray))
	A2 := Article{}
	_ = json.Unmarshal(A1ByteArray, &A2)
	fmt.Println(A2)
	//encode function create a json representation of a struct and writes to the http response
	_ = json.NewEncoder(w).Encode(A2)

}
