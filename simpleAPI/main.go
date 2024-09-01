package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	Articles = append(Articles, Article{ID: "4", Title: "test title3", Description: "test description3", Content: "test content3"})

	HandleRequest()
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", ShowHomePage).Methods("GET")
	myRouter.HandleFunc("/articles", ShowArticles).Methods("GET")
	myRouter.HandleFunc("/articles", AddNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", ShowArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
	CreateToken("testusername", "testusersid", "testuserid")
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

func CreateToken(username string, usersid string, userId string) {
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	Claims := Claims{Username: username, UsersId: usersid, UserId: userId, Expires: expirationTime}
	ClaimsJson, err := json.Marshal(Claims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ClaimsJson)
	header := `{"alg": "HS256", "typ": "JWT"}`
	headerJson := []byte(header)
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJson)
	claimsEncoded := base64.RawURLEncoding.EncodeToString(ClaimsJson)
	fmt.Println("this")
	fmt.Println(headerEncoded)
	fmt.Println(claimsEncoded)

}
