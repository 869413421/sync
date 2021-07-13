package tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sync/bootstarp"
	"testing"
)

func TestLogin(test *testing.T) {
	url := "/login"
	params := make(map[string]string)
	params["username"] = "admin"
	params["password"] = "12345678"

	response := PostJson(url, params, bootstarp.Router)

	result := response.Result()
	assert.Equal(test, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(test, err, "Error Not Nil")

	data := responseData["data"]
	assert.NotNil(test, data, "Response Data Nil")
}
