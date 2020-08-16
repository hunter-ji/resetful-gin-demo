package user

import "github.com/gin-gonic/gin"

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":   "Kuari",
		"site":   "https://justmylife.cc",
		"github": "https://github.com/Kuari",
	})
}
