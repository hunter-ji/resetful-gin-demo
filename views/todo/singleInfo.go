package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func SingleInfo(c *gin.Context) {
	type SingleTodo struct {
		Id int `binding:"required"`
	}

	var singleTodo SingleTodo
	if c.ShouldBind(&singleTodo) != nil {
		c.JSON(200, gin.H{
			"code":    40000,
			"message": "参数不全",
		})
		return
	}

	var todo models.Todo
	db := models.DBConnect()
	if db.Where("id = ?", singleTodo.Id).First(&todo).Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "无效数据",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"id":         todo.ID,
			"title":      todo.Title,
			"created_at": todo.CreatedAt.Format("2006-01-02"),
			"updated_at": todo.UpdatedAt.Format("2006-01-02"),
		},
	})
}
