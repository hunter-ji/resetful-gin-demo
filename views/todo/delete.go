package todo

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Delete(c *gin.Context) {
	type ToDo struct {
		TodoId int `binding:"required" json:"todo_id"`
	}

	var todo ToDo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"message": "参数不全",
		})
		return
	}

	db := models.DBConnect()
	tx := db.MustBegin()
	tx.MustExec("update todo set is_deleted = 1 where todo_id = ?", todo.TodoId)
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 20000,
	})
}
