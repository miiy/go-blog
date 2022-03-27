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
