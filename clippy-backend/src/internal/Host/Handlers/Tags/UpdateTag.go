package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Models "backend/internal/Infrastructure/Models"
	Redis "backend/internal/Infrastructure/Redis"
	Services "backend/internal/Services"
	Utils "backend/internal/Services/Utils"
)

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateTag")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var updatedTag Models.Tag

	rh := Redis.ConnectToRedis()
	t := Utils.ParseRequestBody(r) 
	if t, ok := t.(Models.Tag); ok {
		updatedTag = t
	}

	fmt.Println(updatedTag)

	tags := Services.UpdateTag(rh, updatedTag)

	jsonBytes, err := json.Marshal(tags)
	if err != nil {
		fmt.Printf("Error")
	}
	w.Write(jsonBytes)
}
