package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miiy/go-blog/pkg/gin/middleware"
	"github.com/miiy/go-blog/web"
	"html/template"
	"io/fs"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	// fs
	resourcesFS := web.FS

	// template
	t, err := template.ParseFS(resourcesFS, "resources/templates/*/*.tmpl")
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	// statics
	staticsFs, _ := fs.Sub(resourcesFS, "statics")
	r.StaticFS("/statics", http.FS(staticsFs))

	// favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		file, _ := resourcesFS.ReadFile("resources/statics/favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", file)
	})

	// uploads
	r.Static("/uploads", "./storage/uploads")


	// index
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home/index.tmpl",  gin.H{
			"PageTitle": "Home",
			"Content": "Hello, world.",
		})
	})

	// book
	r.GET("/books", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "books",
		})
	})
	// book detail
	r.GET("/book/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "books detial",
		})
	})

	// middleware
	r.Use(middleware.RequestInfo())

}