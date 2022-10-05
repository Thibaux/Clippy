package main

import (
	"log"
	router "backend/internal/Host"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	router.Router()
}
