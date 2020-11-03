package todo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/gotest"
	"resetful-gin-demo/routers"
)

func TestTodoCreate(t *testing.T) {

	token := gotest.GenToken()

	method := "POST"
	urlStr := "/todo?token=" + token
	data := url.Values{}
	data.Add("Title", "hello")

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/json")
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
