package gotest

import (
	"context"
	"fmt"
	"os"
	"time"

	"resetful-gin-demo/models"
	"resetful-gin-demo/utils"
)

func GenToken() (token string) {
	type User struct {
		ID       int
		Username string `binding:"required"`
		Password string `binding:"required"`
	}

	var user User
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	db := models.DBConnect()
	result := db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		panic("查询用户失败")
	}

	token = utils.Md5(fmt.Sprintf("%s==>%s==>%s", user.Username, user.Password, time.Now()))

	rdb := utils.RedisConnect()
	rdb.HSet(context.Background(), token, "name", user.Username)
	rdb.HSet(context.Background(), token, "uid", user.ID)
	rdb.Expire(context.Background(), token, time.Duration(24*time.Hour))
	return
}
