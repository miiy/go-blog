package home

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/", indexHandler)
}

func Templates() map[string][]string {
	return map[string][]string{
		"home/index": {"layout/layout.html", "layout/header.html", "layout/footer.html", "home/index.html"},
	}
}