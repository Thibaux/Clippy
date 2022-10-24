package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	HealthHandler "backend/internal/Host/Handlers/Health"
	ItemsHandler "backend/internal/Host/Handlers/Items"
	TagsHandlers "backend/internal/Host/Handlers/Tags"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HealthHandler.GetHealth).Methods("GET")
	router.HandleFunc("/health", HealthHandler.GetHealth).Methods("GET")
	router.HandleFunc("/items", ItemsHandler.GetItems).Methods("GET")
	router.HandleFunc("/items", ItemsHandler.CreateItem).Methods("POST")
	router.HandleFunc("/tags", TagsHandlers.GetTags).Methods("GET")
	router.HandleFunc("/tags", TagsHandlers.UpdateTag).Methods("POST")

	var port = os.Getenv("PORT")

	fmt.Printf("Server starting on port: %v", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
