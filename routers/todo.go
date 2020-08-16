package routers

import (
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/views/todo"
)

func LoadTodo(e *gin.Engine) {
	e.GET("/todo", todo.Read)
	e.POST("/todo", todo.Create)
	e.PUT("/todo", todo.Change)
	e.DELETE("/todo", todo.Delete)
}
