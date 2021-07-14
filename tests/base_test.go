package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"sync/bootstarp"
	"sync/pkg/types"
)

func init() {
	bootstarp.Run()
}

var Token = ""

type ResponseData struct {
	Code     int64       `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data"`
}

func ParseResponse(input []byte) (*ResponseData, error) {
	var data = &ResponseData{}
	err := json.Unmarshal(input, data)
	return data, err
}

func Request(method string, url string, router *gin.Engine) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, url, nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func RequestFrom(method string, url string, params map[string]string, router *gin.Engine) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, url, strings.NewReader(types.StrMapToString(params)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
