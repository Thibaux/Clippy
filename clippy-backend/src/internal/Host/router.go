package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	Handlers "backend/internal/Host/Handlers"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Handlers.GetHealth).Methods("GET")
	router.HandleFunc("/health", Handlers.GetHealth).Methods("GET")
	router.HandleFunc("/items", Handlers.GetItems).Methods("GET")
	router.HandleFunc("/items", Handlers.CreateItem).Methods("POST")

	var port = os.Getenv("PORT")

	fmt.Printf("Server starting on port: %v", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
