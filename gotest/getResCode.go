package gotest

import "encoding/json"

func GetResCode(response string) (code int) {
	type Res struct {
		Code int
	}

	var res Res
	json.Unmarshal([]byte(response), &res)

	return res.Code
}
