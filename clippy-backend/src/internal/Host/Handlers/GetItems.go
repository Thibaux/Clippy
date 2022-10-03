package Handlers

import (
	"fmt"
	"net/http"

	Services "backend/internal/Services"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getArticles")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	w.Header().Set("Content-Type", "application/json")
	// jsonBytes, err := json.Marshal(Articles)
	// if err != nil {
	// 	fmt.Printf("Error")
	// }

	Services.GetItems()

	// if err != nil {
	// 	fmt.Printf("Noo")
	// }

	// fmt.Printf(result)
	
	w.WriteHeader(http.StatusOK)
	// w.Write(jsonBytes)
}
