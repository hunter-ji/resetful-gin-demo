package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"resetful-gin-demo/routers"
)

type TestConfig struct {
	urlStr string
}

func NewTest(urlStr string) TestConfig {
	var this TestConfig
	this.urlStr = urlStr
	return this
}

func (tc *TestConfig) Get() *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", tc.urlStr, nil)
	w := httptest.NewRecorder()
	router := routers.SetupRouter()
	router.ServeHTTP(w, req)
	return w
}

func (tc *TestConfig) OptionJson(method string, body map[string]interface{}) *httptest.ResponseRecorder {
	jsonByte, _ := json.Marshal(body)
	req := httptest.NewRequest(method, tc.urlStr, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router := routers.SetupRouter()
	router.ServeHTTP(w, req)
	return w
}
