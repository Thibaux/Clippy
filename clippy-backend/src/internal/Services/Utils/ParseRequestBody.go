package Utils

import (
	"backend/internal/Infrastructure/Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseRequestBody(r *http.Request) Models.Item {
	var item Models.Item

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	error := json.Unmarshal(body, &item)
	if error != nil {
		log.Fatalf("Failed to JSON Unmarshal")
	}

	return item
}