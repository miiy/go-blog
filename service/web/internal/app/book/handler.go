package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/paginater"
	"strconv"
)

func indexHandler(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	page, _:= strconv.Atoi(c.Param("page"))
	pageSize := 2
	list, err := book.service.ListBooks(cid, page, pageSize)
	if err != nil {

	}
	p := paginater.New(int(list.Total), pageSize, page, 5)
	fmt.Printf("%v", list)
	c.HTML(200, "book/list", gin.H{
		"PageTitle": "Article list",
		"BookList": list,
		"Page": p,
	})
}

func showHandler(c *gin.Context) {
	id, _:= strconv.Atoi(c.Param("id"))
	item, err := book.service.GetBook(id)
	if err != nil {

	}
	c.HTML(200, "book/detail", gin.H{
		"PageTitle": "Article detail",
		"Book": item,
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
