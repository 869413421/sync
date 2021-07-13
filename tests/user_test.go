package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"sync/bootstarp"
	"testing"
)

func TestUser(t *testing.T) {
	url := "/user?token=" + Token
	fmt.Println(url)
	response := Get(url, bootstarp.Router)

	result := response.Result()
	assert.Equal(t, 200, result.StatusCode, "StatusCode Not Equal")
	defer result.Body.Close()

	_body, _ := ioutil.ReadAll(result.Body)
	responseData, err := ParseResponse(_body)
	assert.NoError(t, err, "Error Not Nil")
	fmt.Println(responseData.ErrorMsg)
	data := responseData.Data
	assert.NotNil(t, data, "Response Data Nil")

}
