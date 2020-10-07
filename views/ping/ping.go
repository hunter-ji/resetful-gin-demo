package ping

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
	"resetful-gin-demo/utils"
)

func Ping(c *gin.Context) {
	db := models.DBConnect()
	if db.First(&models.User{}).Error != nil {
		c.String(200, "db error")
		return
	}

	rdb := utils.RedisConnect()
	if rdb.Exists(c, "test").Err() != nil {
		c.String(200, "redis error")
		return
	}

	c.String(200, "pong")
}
