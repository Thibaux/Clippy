package Services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"

	Models "backend/internal/Infrastructure/Models"
	Utils "backend/internal/Services/Utils"
)

func GetAllItems(rh *rejson.Handler) Models.Items {
	var items Models.Items

	res, err := redis.Bytes(rh.JSONGet("items", "."))
	if err != nil {
		fmt.Printf("Failed to JSONGet")
		return items
	}
	Utils.HandleError(err)

	error := json.Unmarshal(res, &items)
	if error != nil {
		log.Fatalf("Failed to JSON Unmarshal")
	}

	return items
}

func CreateOneItem(rh *rejson.Handler, item Models.Item) Models.Items {
	items := GetAllItems(rh)
	items = append(items, item)

	res, err := rh.JSONSet("items", ".", items)
	if err != nil {
		fmt.Printf("Failed to JSONGet")
	}
	fmt.Println(res)
	Utils.HandleError(err)

	return items
}
