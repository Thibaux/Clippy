package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Redis "backend/internal/Infrastructure/Redis"
	Services "backend/internal/Services"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetTags")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rh := Redis.ConnectToRedis()
	tags := Services.GetAllTags(rh)

	jsonBytes, err := json.Marshal(tags)
	if err != nil {
		fmt.Printf("Error")
	}
	w.Write(jsonBytes)
}
