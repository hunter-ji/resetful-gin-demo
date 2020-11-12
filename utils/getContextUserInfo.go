package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContextUserInfo struct {
	Uid      int
	Username string
}

func GetContextUserInfo(c *gin.Context) (ContextUserInfo, error) {
	var contextUserInfo ContextUserInfo
	err := func() error {
		// user_id
		userId, contextUserIdErr := c.Get("user_id")
		userIdInt, userIdIntErr := strconv.Atoi(fmt.Sprint(userId))
		if !contextUserIdErr || userIdIntErr != nil {
			return errors.New("user_id")
		}
		contextUserInfo.Uid = userIdInt

		// username
		userName, contextUserErr := c.Get("username")
		if !contextUserErr {
			return errors.New("username")
		}
		contextUserInfo.Username = fmt.Sprint(userName)

		return nil
	}()
	if err != nil {
		return contextUserInfo, err
	}

	return contextUserInfo, nil
}
