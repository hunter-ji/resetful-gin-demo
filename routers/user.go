package routers

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/views/user"
)

func LoadUser(e *gin.Engine) {
	userRouter := e.Group("/user")
	{
		userRouter.POST("/login", user.Login)
		userRouter.POST("/logout", user.Logout)
		userRouter.GET("/info", user.Info)
	}
}
