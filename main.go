package main

import (
	"resetful-gin-demo/routers"
)

func main() {

	// models.CreateDB()
	r := routers.SetupRouter()
	r.Run()
}
