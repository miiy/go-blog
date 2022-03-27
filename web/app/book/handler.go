package book

import (
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "book/list", gin.H{
		"PageTitle": "Article list",
	})
}

func showHandler(c *gin.Context) {
	c.HTML(200, "book/detail", gin.H{
		"PageTitle": "Article detail",
	})
}

func createHandler(ctx *gin.Context) {
	//response.Success(ctx, "story")
}

func storeHandler(ctx *gin.Context) {
	//response.Success(ctx, "story")
}

func editHandler(ctx *gin.Context) {
	//response.Success(ctx, "edit")
}

func updateHandler(ctx *gin.Context) {
	//response.Success(ctx, "update")
}

func destroyHandler(ctx *gin.Context) {
	//response.Success(ctx, "destroy")
}
