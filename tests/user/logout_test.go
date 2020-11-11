package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/tests"
)

func TestLogout(t *testing.T) {
	token := os.Getenv("TOKEN")
	urlStr := "/user/logout?token=" + token

	res := tests.NewTest(urlStr)
	w := res.Get()

	assert.Equal(t, 200, w.Code)

	var response map[string]int
	json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["code"]

	assert.True(t, exists)
	assert.Equal(t, 20000, value)
}
