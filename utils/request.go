package utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func whiteList(url *url.URL) bool {
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]

	noAuthWhiteList := map[string]bool{
		"/user/login": true,
	}

	if _, ok := noAuthWhiteList[queryUrl]; ok {
		return true
	}
	return false
}

func Authorize() gin.HandlerFunc {
	type QueryToken struct {
		Token string `form:"token" binding:"required"`
	}

	return func(c *gin.Context) {

		// 当路由不在白名单之中时进行token检测
		if !whiteList(c.Request.URL) {
			var queryToken QueryToken
			if c.ShouldBind(&queryToken) != nil {
				fmt.Println(queryToken.Token)
				c.AbortWithStatusJSON(200, gin.H{
					"code":    50008,
					"message": "访问未授权",
				})
				return
			}
			fmt.Println(queryToken.Token)

			rdb := RedisConnect()
			nameValue, nameErr := rdb.HGet(c, queryToken.Token, "name").Result()
			if nameErr != nil {
				if nameErr == redis.Nil {
					c.AbortWithStatusJSON(200, gin.H{
						"code":    50008,
						"message": "无效token",
					})
					return
				}
			}

			userIdValue, userIdErr := rdb.HGet(c, queryToken.Token, "uid").Result()
			if userIdErr != nil {
				if userIdErr == redis.Nil {
					c.AbortWithStatusJSON(200, gin.H{
						"code":    50008,
						"message": "无效token",
					})
					return
				}
			}

			// 上下文变量
			c.Set("name", nameValue)
			c.Set("uid", userIdValue)
			c.Set("token", queryToken.Token)

		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatusJSON(200, gin.H{
				"code": 20000,
			})
			return
		}

		c.Next()

	}
}
