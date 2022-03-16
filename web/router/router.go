package router

import (
	"github.com/gin-contrib/multitemplate"
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
	//t, err := template.ParseFS(resourcesFS, "resources/templates/*/*.html")
	//if err != nil {
	//	panic(err)
	//}
	//r.SetHTMLTemplate(t)
	r.HTMLRender = createMyRender()

	// statics
	staticsFs, _ := fs.Sub(resourcesFS, "resources/statics")
	r.StaticFS("/statics", http.FS(staticsFs))

	// favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		file, _ := resourcesFS.ReadFile("resources/statics/favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", file)
	})

	// uploads
	r.Static("/uploads", "./storage/uploads")

	// pages
	r.GET("/pages/list", func(c *gin.Context) {
		c.HTML(200, "pages/list.html",  gin.H{})

	})
	r.GET("/pages/detail", func(c *gin.Context) {
		c.HTML(200, "pages/detail.html",  gin.H{})
	})

	// index
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index",  gin.H{
			"PageTitle": "Home",
			"Content": "Hello, world.",
			"Header": "header.",
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


func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	resourcesFS := web.FS
	t, err := template.ParseFS(resourcesFS, "resources/templates/layout/*.html", "resources/templates/home/index.html")
	if err != nil {
		panic(err)
	}
	r.Add("index", t)
	return r
}