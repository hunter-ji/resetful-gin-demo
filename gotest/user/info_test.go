package gotest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/gotest"
	"resetful-gin-demo/routers"
)

func TestUserInfo(t *testing.T) {

	token := gotest.GenToken()
	method := "GET"
	urlStr := "/user/info?token=" + token

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlStr, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]int
	json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["code"]

	assert.True(t, exists)
	assert.Equal(t, 20000, value)
}
