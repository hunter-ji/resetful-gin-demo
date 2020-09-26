package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(originMessage string) (result string) {
	data := []byte(originMessage)
	result = fmt.Sprintf("%x", md5.Sum(data))
	return
}
