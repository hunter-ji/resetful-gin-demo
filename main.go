package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/routers"
	"resetful-gin-demo/utils"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(utils.Authorize())
	routers.LoadUser(r)
	routers.LoadTodo(r)
	r.Run()
}
