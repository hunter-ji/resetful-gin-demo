package todo

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Create(c *gin.Context) {
	type ToDo struct {
		Title string `binding:"required,gte=1,lte=30"`
	}

	var todo ToDo
	if c.ShouldBindJSON(&todo) != nil {
		c.JSON(200, gin.H{
			"code":    40000,
			"message": "参数不全",
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

	contextUidInt, contextUidIntErr := strconv.Atoi(fmt.Sprint(contextUid))
	if contextUidIntErr != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "无效请求",
		})
		return
	}

	db := models.DBConnect()
	tx := db.MustBegin()
	res := tx.MustExec("insert into todo (title, user_id) values (?, ?)", todo.Title, contextUidInt)
	insertedTodoId, insertedTodoErr := res.LastInsertId()
	if insertedTodoErr != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "创建失败",
		})
		return
	}

	tx.Commit()

	fmt.Println(insertedTodoId)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]int{
			"id": int(insertedTodoId),
		},
	})
}
