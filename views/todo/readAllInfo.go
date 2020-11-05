package todo

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func ReadAllInfo(c *gin.Context) {
	type Info struct {
		models.User
		models.Todo
	}

	db := models.DBConnect()

	var info []Info
	allInfoSelectErr := db.Select(&info,
		"select user.username, todo.todo_id, todo.title, todo.user_id, todo.created_at "+
			"from (select * from todo where is_deleted = 0) as todo "+
			"inner join (select * from user where is_deleted = 0) as user on user.user_id = todo.user_id")
	if allInfoSelectErr != nil {
		fmt.Println(allInfoSelectErr)
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "查询数据错误",
		})
		return
	}

	var data []map[string]interface{}
	for i := 0; i < len(info); i++ {
		todoObject := map[string]interface{}{
			"username":   info[i].User.Username,
			"todo_id":    info[i].Todo.TodoID,
			"title":      info[i].Todo.Title,
			"user_id":    info[i].User.UserID,
			"created_at": info[i].Todo.CreatedAt.Format("2006-01-02"),
		}
		data = append(data, todoObject)
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": data,
	})
}
