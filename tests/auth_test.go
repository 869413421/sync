package tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sync/bootstarp"
	"testing"
)

func TestLogin(t *testing.T) {
	url := "/login"
	params := make(map[string]string)
	params["username"] = "admin"
	params["password"] = "12345678"

	response := PostJson(url, params, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)

	assert.NoError(t, err, "Error Not Nil")

	data := responseData.Data.(map[string]interface{})

	assert.NotNil(t, data, "Response Data Nil")

	Token = data["token"].(string)
}
