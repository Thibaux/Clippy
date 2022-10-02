package router

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"

	Handlers "backend/internal/Host/Handlers"
)

type Item struct {
	Name string
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getHealth")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "API is healthy"}`))
}

func Router() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", getHealth)
	router.HandleFunc("/health", getHealth)
	router.HandleFunc("/items", Handlers.GetItems)

	var port = ":8080"
	fmt.Printf("Server starting on port: %v", port)
	log.Fatal(http.ListenAndServe(port, router))
}
