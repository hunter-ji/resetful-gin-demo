package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Delete(c *gin.Context) {
	type ToDo struct {
		Id int `binding:"required,lte=0"`
	}

	var todo ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	db := models.DBConnect()
	deleteRes := db.Delete(&models.Todo{}, todo.Id)
	if deleteRes.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "删除失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
