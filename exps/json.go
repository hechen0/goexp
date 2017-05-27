package exps

import (
	"encoding/json"
	"strings"
	"fmt"
	"log"
)

type newJson struct {
	name string `json:a_name`
	age  int `json:a_age`
}

func DecodeJson() {
	a := `[
		{"A_Name": "a name", "A_Age": "a age"},
		{"a_nAme": "b name", "b_AGE": "b age"}
	]`

	var cc []newJson
	err := json.NewDecoder(strings.NewReader(a)).Decode(&cc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", cc)
}
