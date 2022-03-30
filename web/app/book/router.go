package book

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/books", indexHandler)
	r.GET("/books/:id", showHandler)
	r.GET("/books/:id/create", createHandler)
	r.POST("/books/:id", storeHandler)
	r.GET("/books/:id/edit", editHandler)
	r.PUT("/books/:id", updateHandler)
	r.DELETE("/books/:id", destroyHandler)
}

func Templates() map[string][]string {
	return map[string][]string{
		"book/detail": {"layout/layout.html", "layout/header.html", "layout/footer.html", "book/detail.html"},
		"book/list": {"layout/layout.html", "layout/header.html", "layout/footer.html", "book/list.html"},
	}
}