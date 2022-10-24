package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nitishm/go-rejson/v4"

	Models "backend/internal/Infrastructure/Models"
	Redis "backend/internal/Infrastructure/Redis"
	Services "backend/internal/Services"
	Utils "backend/internal/Services/Utils"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateItem")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var item Models.Item

	rh := Redis.ConnectToRedis()
	i := Utils.ParseRequestBody(r)
	if i, ok := i.(Models.Item); ok {
		item = i
	} 

	item = enrichItemData(item, rh)
		
	updatedItems := Services.CreateOneItem(rh, item)
	jsonBytes, err := json.Marshal(updatedItems)
	if err != nil {
		fmt.Printf("Error")
	}

	w.Write(jsonBytes)
}


func enrichItemData(item Models.Item, rh *rejson.Handler) Models.Item {
	currentTime := time.Now()

	item.Id = uuid.New().String()
	item.CreatedAt = currentTime.Format("2006-01-02 15:04:05 Monday")
	item.UpdatedAt = currentTime.Format("2006-01-02 15:04:05 Monday")

	for i := 0; i < len(item.Tags); i++ {
		item.Tags[i] = updateTag(item.Tags[i], rh)
	}
	
	return item
}

func updateTag(tag Models.Tag, rh *rejson.Handler) Models.Tag {
	serverTags := Services.GetAllTags(rh)

	for i := 0; i < len(serverTags); i++ {
		if (serverTags[i].Id == tag.Id) {
			tag.TimesUsed++
			Services.UpdateTag(rh, tag)
		}
	}

	return tag
}