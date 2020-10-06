package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
	"resetful-gin-demo/utils"
)

func Login(c *gin.Context) {
	type User struct {
		ID       int
		Username string `binding:"required"`
		Password string `binding:"required"`
	}

	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	db := models.DBConnect()
	result := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"code":    20001,
			"message": "账户不存在或者密码不对",
		})
		return
	}

	token := utils.Md5(fmt.Sprintf("%s==>%s==>%s", user.Username, user.Password, time.Now()))

	rdb := utils.RedisConnect()
	rdb.HSet(c, token, "name", user.Username)
	rdb.HSet(c, token, "uid", user.ID)
	rdb.Expire(c, token, time.Duration(24*time.Hour))

	c.JSON(200, gin.H{
		"code": 20000,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
