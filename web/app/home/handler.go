package home

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "home/index", gin.H{
		"PageTitle": "Home",
		"Content":   "Hello, world.",
		"Header":    "header.",
	})
}
