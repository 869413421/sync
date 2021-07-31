package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sync/bootstarp"
	"sync/pkg/types"
	"testing"
)

var UserID int64

func TestUser(t *testing.T) {
	url := "/user?token=" + Token
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

func TestCreateUser(t *testing.T) {
	url := "/user?token=" + Token
	params := make(map[string]string)
	params["name"] = "test"
	params["password"] = "123456"
	params["email"] = "test@163.com"
	params["avatar"] = "www.baidu.com"

	response := RequestFrom("POST", url, params, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data.(map[string]interface{})
	assert.NotNil(t, data, "Response Data Nil")
	UserID = int64(data["ID"].(float64))
}

func TestGetUser(t *testing.T) {
	url := "/user/" + types.Int64ToString(UserID) + "?token=" + Token
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

func TestUpdateUser(t *testing.T) {
	url := "/user/" + types.Int64ToString(UserID) + "?token=" + Token
	params := make(map[string]string)
	params["name"] = "testupdate"
	params["password"] = "123456"
	params["email"] = "testupdate@163.com"
	params["avatar"] = "www.baidu.com"

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

func TestDeleteUser(t *testing.T) {
	url := "/user/" + types.Int64ToString(UserID) + "?token=" + Token
	response := Request("DELETE", url, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	data := responseData.Data
	assert.NotNil(t, data, "Response Data Nil")
	fmt.Println(data)
}
