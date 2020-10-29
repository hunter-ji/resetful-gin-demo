package todo

import (
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
	data.Set("title", `hello`)

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic("请求测试失败")
	}
	router.ServeHTTP(w, req)

	resCode := gotest.GetResCode(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 20000, resCode)
}
