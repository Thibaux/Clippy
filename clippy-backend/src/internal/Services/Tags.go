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

func GetAllTags(rh *rejson.Handler) Models.Tags {
	var tags Models.Tags

	res, err := redis.Bytes(rh.JSONGet("tags", "."))
	if err != nil {
		fmt.Printf("Failed to JSONGet")
		return tags
	}
	Utils.HandleError(err)

	error := json.Unmarshal(res, &tags)
	if error != nil {
		log.Fatalf("Failed to JSON Unmarshal")
	}

	return tags
}


func UpdateTag(rh *rejson.Handler, updatedTag Models.Tag) Models.Tags {
	tags := GetAllTags(rh)

	fmt.Println(updatedTag)

	for i := 0; i < len(tags); i++ {
		if (tags[i].Id == updatedTag.Id) {

			fmt.Println(updatedTag)

			tags[i].Id = updatedTag.Id
			tags[i].Name = updatedTag.Name
		}
	}

	res, err := rh.JSONSet("tags", ".", tags)
	if err != nil {
		fmt.Printf("Failed to JSONSet")
	}
	fmt.Println(res)
	Utils.HandleError(err)

	return tags
}

