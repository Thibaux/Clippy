package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Redis "backend/internal/Infrastructure/Redis"
	Services "backend/internal/Services"
	Utils "backend/internal/Services/Utils"
	Models "backend/internal/Infrastructure/Models"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item Models.Item

	fmt.Println("Endpoint Hit: CreateItem")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rh := Redis.ConnectToRedis()
	item = Utils.ParseRequestBody(r)

    id := "1234234"
    item = append(item, id)

	updatedItems := Services.CreateOneItem(rh, item)
	jsonBytes, err := json.Marshal(updatedItems)
	if err != nil {
		fmt.Printf("Error")
	}

	w.Write(jsonBytes)
}
