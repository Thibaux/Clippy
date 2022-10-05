package Handlers

import (
	"fmt"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetHealth")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "API is healthy"}`))
}
