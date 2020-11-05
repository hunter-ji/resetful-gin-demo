package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Change(c *gin.Context) {
	type ToDo struct {
		TodoId int    `binding:"required"`
		Title  string `binding:"gte=0,lte=30"`
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
	db := models.DBConnect()
	tx := db.MustBegin()
	tx.MustExec("update todo set title = ? where todo_id = ?", todo.Title, todo.TodoId)
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
