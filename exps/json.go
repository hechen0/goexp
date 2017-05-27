package exps

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Language struct {
	Id   int `json:"a_id"`
	Name string `json:"a_name"`
}

func DecodeJson() {

	//String contains two JSON rows.
	text := `
	[
	{"a_id": 100, "a_name": "Go"},{"A_iD": 200, "A_NaMe": "Java"
	}
	]
	`
	//text := "[{\"Id\": 100, \"Name\": \"Go\"}, {\"Id\": 200, \"Name\": \"Java\"}]"
	// Get byte slice from string.

	// Unmarshal string into structs.
	var languages []Language
	err := json.NewDecoder(strings.NewReader(text)).Decode(&languages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", languages)
}
