package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"PageTitle": "Home",
		"Content": "Hello, world.",
	})
}
