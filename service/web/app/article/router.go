package article

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/articles", indexHandler)
	r.GET("/articles/:id", showHandler)
	r.GET("/articles/:id/create", createHandler)
	r.POST("/articles/:id", storeHandler)
	r.GET("/articles/:id/edit", editHandler)
	r.PUT("/articles/:id", updateHandler)
	r.DELETE("/articles/:id", destroyHandler)
}

func Templates() map[string][]string {
	return map[string][]string{
		"article/detail": {"layout/layout.html", "layout/header.html", "layout/footer.html", "article/detail.html"},
		"article/list":   {"layout/layout.html", "layout/header.html", "layout/footer.html", "article/list.html"},
	}
}
