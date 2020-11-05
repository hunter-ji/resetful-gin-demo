package todo

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Read(c *gin.Context) {

	contextUid, contextUidExits := c.Get("uid")
	if !contextUidExits {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "无效请求",
		})
		return
	}

	db := models.DBConnect()

	var todoList []models.Todo
	todoListSelectErr := db.Select(&todoList,
		"select * from todo where user_id = ? and is_deleted = 0 order by todo_id desc", contextUid)
	if todoListSelectErr != nil {
		fmt.Println(todoListSelectErr)
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "查询数据错误",
		})
		return
	}

	var data []map[string]interface{}
	for i := 0; i < len(todoList); i++ {
		todoObject := map[string]interface{}{
			"id":         todoList[i].TodoID,
			"title":      todoList[i].Title,
			"created_at": todoList[i].CreatedAt.Format("2006-01-02"),
		}
		data = append(data, todoObject)
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": data,
	})
}
