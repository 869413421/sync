package tests

import (
	"encoding/json"
	"sync/bootstarp"
)

func init() {
	bootstarp.Run()
}

var BaseUrl = "http://localhost:8989"

func ParseResponse(input string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	return data, err
}
