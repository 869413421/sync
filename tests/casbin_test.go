package tests

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sync/bootstarp"
	"sync/pkg/types"
	"testing"
)

var RuleId int64

func TestRule(t *testing.T) {
	url := "/casbin??ptype=g&token=" + Token
	response := Request("GET", url, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")

	data := responseData.Data
	assert.NotNil(t, data, "Response Data Nil")
}

func TestCreateRule(t *testing.T) {
	url := "/casbin?token=" + Token
	params := make(map[string]string)
	params["ptype"] = "g"
	params["v0"] = "test"
	params["v1"] = "test"
	params["v2"] = "test"
	params["v3"] = "test"

	response := RequestFrom("POST", url, params, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data.(map[string]interface{})
	assert.NotNil(t, data, "Response Data Nil")
	RuleId = int64(data["ID"].(float64))
}

func TestGetRule(t *testing.T) {
	url := "/casbin/" + types.Int64ToString(RuleId) + "?token=" + Token
	response := Request("GET", url, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data.(map[string]interface{})
	assert.NotNil(t, data, "Response Data Nil")
}

func TestUpdateRule(t *testing.T) {
	url := "/casbin/" + types.Int64ToString(RuleId) + "?token=" + Token
	params := make(map[string]string)
	params["ptype"] = "g"
	params["v0"] = "update"
	params["v1"] = "update"
	params["v2"] = "update"
	params["v3"] = "update"

	response := RequestFrom("PUT", url, params, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data.(map[string]interface{})
	assert.NotNil(t, data, "Response Data Nil")
}

func TestDeleteRule(t *testing.T) {
	url := "/casbin/" + types.Int64ToString(RuleId) + "?token=" + Token
	response := Request("DELETE", url, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data
	assert.NotNil(t, data, "Response Data Nil")
}
