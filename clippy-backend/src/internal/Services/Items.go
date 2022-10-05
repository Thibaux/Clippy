package Services

import (
	"fmt"

	"github.com/nitishm/go-rejson/v4"

	Models	"backend/internal/Infrastructure/Models"
	// Redis 	"backend/internal/Infrastructure/Redis"
	// Utils 	"backend/internal/Services/Utils"
)

type Item struct {
	id		string
	name	string
}

func GetAllItems(rh *rejson.Handler) Models.Items {
	fmt.Println("hsdf")

	var items Models.Items

	i := Models.Item{Id: "123", Name: "thisName"}

	items = append(items, i)

	Ii := Models.Item

	res, err := rh.JSONSet("items", ".", Ii)
	if err != nil {
		fmt.Printf("Failed to JSONGet")
		return items
	}
	fmt.Println("obj:", res)

	// res, err := rh.JSONGet("items", ".")
	// if err != nil {
	// 	fmt.Printf("Failed to JSONGet")
	// 	return items
	// }

	fmt.Println("got obj:", res)

	// Utils.HandleError(err)
	

	return items
}