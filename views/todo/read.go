package todo

import (
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 20000,
		"data": []map[string]interface{}{
			{
				"todoId":  1,
				"title":   "Hello, World !",
				"addTime": "2020-08-16",
			},
			{
				"todoId":  2,
				"title":   "Hello, World2 !",
				"addTime": "2020-08-16",
			},
			{
				"todoId":  3,
				"title":   "Hello, World3 !",
				"addTime": "2020-08-16",
			},
		},
	})
}
