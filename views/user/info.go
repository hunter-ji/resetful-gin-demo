package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
)

func Info(c *gin.Context) {

	userId, userIdErr := c.Get("uid")
	if !userIdErr {
		c.JSON(200, gin.H{
			"code":    40001,
			"message": "无效请求",
		})
		return
	}

	var user models.User
	db := models.DBConnect()
	userSelectErr := db.Get(&user, "select * from user where user_id=? limit 1", userId)
	if userSelectErr != nil {
		fmt.Println(userSelectErr)
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
			"created_at":   user.CreatedAt.Format("2006-01-02 12:15:15"),
		},
	})
}
