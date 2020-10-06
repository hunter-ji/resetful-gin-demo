package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/utils"
)

func Logout(c *gin.Context) {
	tokenVal, tokenErr := c.Get("token")
	if !tokenErr {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "无法获取请求用户信息",
		})
		return
	}

	if utils.RedisConnect().Del(c, fmt.Sprintf("%v", tokenVal)).Err() != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "注销用户失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
