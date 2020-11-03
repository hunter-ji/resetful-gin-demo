package main

import (
	"fmt"

	"resetful-gin-demo/routers"
	"resetful-gin-demo/utils"
)

func main() {

	if err := utils.TransInit("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	// models.CreateDB()
	r := routers.SetupRouter()
	r.Run()
}
