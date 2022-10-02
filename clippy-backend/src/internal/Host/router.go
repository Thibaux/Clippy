package router

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/joho/godotenv"

	Handlers "backend/internal/Host/Handlers"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Handlers.GetHealth)
	router.HandleFunc("/health", Handlers.GetHealth)
	router.HandleFunc("/items", Handlers.GetItems)

	var port = os.Getenv("PORT")

	fmt.Printf("Server starting on port: %v", port)
	log.Fatal(http.ListenAndServe(port, router))
}
