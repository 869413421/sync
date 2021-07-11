package tests

import (
	"encoding/json"
)

var BaseUrl = "http://localhost:8989"

func ParseResponse(input string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	return data, err
}
