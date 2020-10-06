package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Info(c *gin.Context) {

	type User struct {
		Id          int
		Username    string
		PhoneNumber string
		CreatedAt   time.Time
	}

	userId, userIdErr := c.Get("uid")
	if !userIdErr {
		c.JSON(200, gin.H{
			"code":    40001,
			"message": "无效请求",
		})
		return
	}

	var user User
	db := models.DBConnect()
	result := db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "账户不存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"username":     user.Username,
			"phone_number": user.PhoneNumber,
			"created_at":   user.CreatedAt.Format("2006-01-02"),
		},
	})
}
