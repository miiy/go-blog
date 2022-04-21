package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func indexHandler(c *gin.Context) {
	page, _:= strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	list, err := book.service.ListBooks(page, pageSize)
	if err != nil {

	}
	fmt.Printf("%v", list)
	c.HTML(200, "book/list", gin.H{
		"PageTitle": "Article list",
		"BookList": list,
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
