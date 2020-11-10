package middleware

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/utils"
)

func whiteList() map[string]string {
	// 白名单
	return map[string]string{
		"/user/login": "POST",
		"/ping":       "GET",
	}
}

func withinWhiteList(url *url.URL, method string) bool {
	noAuthWhiteList := whiteList()
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]

	if _, ok := noAuthWhiteList[queryUrl]; ok {
		if noAuthWhiteList[queryUrl] == method {
			return true
		}
		return false
	}
	return false
}

func Authorize() gin.HandlerFunc {
	type QueryToken struct {
		Token string `form:"token" binding:"required"`
	}

	return func(c *gin.Context) {

		// 当路由不在白名单之中时进行token检测
		if !withinWhiteList(c.Request.URL, c.Request.Method) {
			var queryToken QueryToken
			if c.ShouldBindQuery(&queryToken) != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"code":    50008,
					"message": "访问未授权",
				})
				return
			}

			// 获取用户名和用户id
			rdb := utils.RedisConnect()
			redisInfoRes := rdb.HMGet(c, queryToken.Token, "name", "uid")
			if redisInfoRes.Err() != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"code":    50008,
					"message": "访问未授权",
				})
				return
			}

			for n, _ := range redisInfoRes.Val() {
				if redisInfoRes.Val()[n] == nil {
					c.AbortWithStatusJSON(200, gin.H{
						"code":    50008,
						"message": "访问未授权",
					})
					return
				}
			}

			// 上下文变量
			c.Set("name", fmt.Sprintln(redisInfoRes.Val()[0]))
			c.Set("uid", fmt.Sprintln(redisInfoRes.Val()[1]))
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
