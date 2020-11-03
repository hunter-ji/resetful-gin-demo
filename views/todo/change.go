package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Change(c *gin.Context) {
	type ToDo struct {
		Id    int `binding:"required"`
		Title string `binding:"gte=0, lte=30"`
	}

	var todo ToDo
	if c.ShouldBindJSON(&todo) != nil {
		c.JSON(200, gin.H{
			"code":    40000,
			"message": "参数不全",
		})
		return
	}

	// 更新数据库
	var dbTodo models.Todo
	db := models.DBConnect()
	if db.Where("id = ?", todo.Id).First(&dbTodo).Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "更新失败",
		})
		return
	}
	dbTodo.Title = todo.Title
	db.Save(&dbTodo)

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
