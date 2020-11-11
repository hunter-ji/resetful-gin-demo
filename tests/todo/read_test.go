package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"resetful-gin-demo/tests"
)

func TestTodoRead(t *testing.T) {

	token := os.Getenv("TOKEN")
	urlStr := "/todo?token=" + token

	res := tests.NewTest(urlStr)
	w := res.Get()

	assert.Equal(t, 200, w.Code)

	var response map[string]int
	json.Unmarshal([]byte(w.Body.String()), &response)

	value, exits := response["code"]
	assert.True(t, exits)
	assert.Equal(t, 20000, value)
}
