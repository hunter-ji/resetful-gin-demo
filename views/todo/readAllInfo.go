package todo

import (
	"time"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func ReadAllInfo(c *gin.Context) {
	type Info struct {
		Username  string
		TodoId    int
		Title     string
		UserId    int
		CreatedAt time.Time
	}

	db := models.DBConnect()

	var info []Info
	result := db.Table("todos").
		Order("todos.id desc").
		Select("users.username, todos.id as todo_id, todos.title, todos.user_id, todos.created_at").
		Joins("join users on users.id = todos.user_id").
		Find(&info)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "查询数据错误",
		})
		return
	}

	var data []map[string]interface{}
	for i := 0; i < len(info); i++ {
		todoObject := map[string]interface{}{
			"username":   info[i].Username,
			"id":         info[i].TodoId,
			"title":      info[i].Title,
			"user_id":    info[i].UserId,
			"created_at": info[i].CreatedAt.Format("2006-01-02"),
		}
		data = append(data, todoObject)
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": data,
	})
}
