package main

import (
	"github.com/gin-gonic/gin"
	"goblog.com/web/router"
)

func main()  {
	r := gin.Default()
	router.RegisterRouter(r)
	gin.SetMode(gin.DebugMode)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}