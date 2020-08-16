package todo

import "github.com/gin-gonic/gin"

func Delete(c *gin.Context) {
	type ToDo struct {
		TodoId int `binding:"required"`
	}

	var todo ToDo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
