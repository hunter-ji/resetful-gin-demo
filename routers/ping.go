package routers

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/views/ping"
)

func LoadPing(e *gin.Engine) {
	e.GET("/ping", ping.Ping)
}
