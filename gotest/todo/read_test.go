package todo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/gotest"
	"resetful-gin-demo/routers"
)

func TestTodoRead(t *testing.T) {

	token := gotest.GenToken()

	method := "GET"
	urlStr := "/todo?token=" + token

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		panic("请求测试失败")
	}
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]int
	json.Unmarshal([]byte(w.Body.String()), &response)

	value, exits := response["code"]
	assert.True(t, exits)
	assert.Equal(t, 20000, value)
}
