package tests

import (
	"net/http"
	"testing"
)

func TestUser(t *testing.T) {
	url := BaseUrl + "/user"

	response, err := http.Get(url)
	if err == nil {
		defer response.Body.Close()
	}

	//assert.NoError(test, err, "error not nil")
	//
	////_body, _ := io.ReadAll(response.Body)
	////fmt.Println(ParseResponse(string(_body)))
	//assert.Equal(test, 200, response.StatusCode, "StatusCode Not Equal")
}
