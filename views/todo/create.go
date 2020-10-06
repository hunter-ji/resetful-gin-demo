package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Create(c *gin.Context) {
	type ToDo struct {
		Title string `binding:"required"`
	}

	var todo ToDo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	contextUid, contextUidExits := c.Get("uid")
	if !contextUidExits {
		c.JSON(200, gin.H{
			"code":    40001,
			"message": "无效请求",
		})
	}

	db := models.DBConnect()
	newToDo := models.Todo{Title: todo.Title, UserId: contextUid.(int)}
	insertRes := db.Create(&newToDo)
	if insertRes.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "创建失败",
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
