package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/tests"
)

func TestTodoCreate(t *testing.T) {

	token := os.Getenv("TOKEN")

	method := "POST"
	urlStr := "/todo?token=" + token

	body := map[string]interface{}{
		"title": "hello",
	}

	res := tests.NewTest(urlStr)
	w := res.OptionJson(method, body)

	assert.Equal(t, 200, w.Code)

	var response map[string]int
	json.Unmarshal([]byte(w.Body.String()), &response)

	value, exits := response["code"]
	assert.True(t, exits)
	assert.Equal(t, 20000, value)
}
