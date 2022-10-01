package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



type Item struct {
	Name string
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article


func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

    fmt.Fprint(w, Articles)
    fmt.Println("Endpoint Hit: getArticles")
}

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
    router.HandleFunc("/a", getArticles)
	log.Fatal(http.ListenAndServe(":8080", router))
}