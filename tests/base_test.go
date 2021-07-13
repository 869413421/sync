package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sync/bootstarp"
	"sync/pkg/types"
)

func init() {
	bootstarp.Run()
}

var Token = ""

type ResponseData struct {
	Code     int64                  `json:"code"`
	ErrorMsg string                 `json:"errorMsg"`
	Data     map[string]interface{} `json:"data"`
}

func ParseResponse(input []byte) (*ResponseData, error) {
	var data = &ResponseData{}
	err := json.Unmarshal(input, data)
	return data, err
}

func Get(url string, router *gin.Engine) *httptest.ResponseRecorder {
	request := httptest.NewRequest("GET", url, nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func PostFrom(url string, params map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	request := httptest.NewRequest("POST", url+types.StrMapToString(params), nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func PostJson(url string, params map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(params)

	request := httptest.NewRequest("POST", url, bytes.NewReader(jsonByte))
	request.Header.Add("Content-Type", "application/json")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}
