package Utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseRequestBody(r *http.Request) any {
	var incommingBody any

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	error := json.Unmarshal(body, &incommingBody)
	if error != nil {
		log.Fatalf("Failed to JSON Unmarshal")
	}

	return incommingBody
}