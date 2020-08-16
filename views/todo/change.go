package todo

import "github.com/gin-gonic/gin"

func Change(c *gin.Context) {
	type ToDo struct {
		TodoId int `binding:"required"`
		Title  string
		Name   string
	}

	var todo ToDo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
