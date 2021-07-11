package tests

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogin(test *testing.T) {
	url := BaseUrl + "/login"
	requestBody := fmt.Sprintf(`{
		"username":"%s",
		"password":"%s"
	}`, "admin", "12345678")


	requestJson := []byte(requestBody)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestJson))
	if err == nil {
		defer response.Body.Close()
	}

	assert.NoError(test, err, "error not nil")

	//_body, _ := io.ReadAll(response.Body)
	//fmt.Println(ParseResponse(string(_body)))
	assert.Equal(test, 200, response.StatusCode, "StatusCode Not Equal")
}
