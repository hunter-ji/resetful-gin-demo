package ping

import (
	"context"

	"github.com/gin-gonic/gin"
	"resetful-gin-demo/models"
	"resetful-gin-demo/utils"
)

func Ping(c *gin.Context) {
	db := models.DBConnect()
	if db.Ping() != nil {
		c.String(200, "db error")
		return
	}

	rdb := utils.RedisConnect()
	if rdb.Ping(context.Background()).Err() != nil {
		c.String(200, "redis error")
		return
	}

	c.String(200, "pong")
}
