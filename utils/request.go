package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
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
				c.Abort()
				c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			}
		}

		c.Next()

	}
}
