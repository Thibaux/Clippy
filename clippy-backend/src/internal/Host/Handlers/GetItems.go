package Handlers

import (
	"fmt"
	"net/http"

	"github.com/nitishm/go-rejson/v4"

	Redis	 "backend/internal/Infrastructure/Redis"		
	Services "backend/internal/Services"
)


func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getArticles")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// jsonBytes, err := json.Marshal(Articles)
	// if err != nil {
	// 	fmt.Printf("Error")
	// }

	rh := rejson.NewReJSONHandler()

	cli := Redis.ConnectToRedis()
	rh.SetGoRedisClient(cli)

	items := Services.GetAllItems(rh)
	fmt.Println("v*", items)

	// if err != nil {
	// 	fmt.Printf("Noo")
	// }

	// fmt.Printf(result)

	// w.Write(jsonBytes)
}
