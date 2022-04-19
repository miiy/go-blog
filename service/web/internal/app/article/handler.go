package article

import (
	"github.com/gin-gonic/gin"
	"log"
)

func indexHandler(c *gin.Context) {
	list, err := article.service.ArticleList()
	if err != nil {

	}
	log.Println(list)
	c.HTML(200, "article/list", gin.H{
		"PageTitle": "Article list",
	})
}

func showHandler(c *gin.Context) {
	c.HTML(200, "article/detail", gin.H{
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
