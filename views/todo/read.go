package todo

import (
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

	var todoList []models.Todo
	db := models.DBConnect()
	result := db.Where("user_id = ?", contextUid).Find(&todoList)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "错误数据错误",
		})
		return
	}

	var data []map[string]interface{}
	for i := 0; i < len(todoList); i++ {
		todoObject := map[string]interface{}{
			"id":         todoList[i].ID,
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
