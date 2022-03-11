package about

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about/index.tmpl", gin.H{
		"PageTitle": "About",
		"Content": "about us.",
	})
}
