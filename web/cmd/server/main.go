package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miiy/go-blog/web/router"
)

func main()  {
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}