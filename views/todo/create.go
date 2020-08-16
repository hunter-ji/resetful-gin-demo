package todo

import "github.com/gin-gonic/gin"

func Create(c *gin.Context) {
	type ToDo struct {
		Title string `binding: "required"`
		Name  string `binding: "required"`
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
		"data": map[string]int{
			"todoId": 2,
		},
	})
}
