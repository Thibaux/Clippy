package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Handlers "backend/internal/Host/Handlers"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Handlers.GetHealth)
	router.HandleFunc("/health", Handlers.GetHealth)
	router.HandleFunc("/items", Handlers.GetItems)

	var port = ":8080"
	fmt.Printf("Server starting on port: %v", port)
	log.Fatal(http.ListenAndServe(port, router))
}
