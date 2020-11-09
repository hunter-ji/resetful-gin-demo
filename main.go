package main

import (
	"resetful-gin-demo/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run()
}
