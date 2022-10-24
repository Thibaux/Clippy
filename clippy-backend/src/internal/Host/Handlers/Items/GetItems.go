package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Redis "backend/internal/Infrastructure/Redis"
	Services "backend/internal/Services"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetItems")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rh := Redis.ConnectToRedis()
	items := Services.GetAllItems(rh)

	jsonBytes, err := json.Marshal(items)
	if err != nil {
		fmt.Printf("Error")
	}
	w.Write(jsonBytes)
}
