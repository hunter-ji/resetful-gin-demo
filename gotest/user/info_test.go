package gotest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/gotest"
	"resetful-gin-demo/routers"
)

func TestUserInfo(t *testing.T) {

	token := gotest.GenToken()

	router := routers.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/info?token="+token, nil)
	router.ServeHTTP(w, req)

	resCode := gotest.GetResCode(w.Body.String())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 20000, resCode)
}
