package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/miiy/go-blog/pkg/gin/middleware"
	"github.com/miiy/go-blog/web/resources/assets"
	"github.com/miiy/go-blog/web/resources/templates"
	"html/template"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {

	// template
	tr, err := createTemplateRender()
	if err != nil {
		panic(err)
	}
	r.HTMLRender = tr

	// assets
	r.StaticFS("/assets", http.FS(assets.FS))

	// favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		file, _ := assets.FS.ReadFile("favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", file)
	})

	// uploads
	r.Static("/uploads", "./storage/uploads")

	// pages
	r.GET("/pages/list", func(c *gin.Context) {
		c.HTML(200, "pages/list",  gin.H{})

	})
	r.GET("/pages/detail", func(c *gin.Context) {
		c.HTML(200, "pages/detail",  gin.H{})
	})

	// index
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home/index",  gin.H{
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

func createTemplateRender() (multitemplate.Renderer, error) {
	r := multitemplate.NewRenderer()

	layouts := []string{"layout/layout.html", "layout/header.html", "layout/footer.html"}
	templatesMap := map[string][]string{
		"home/index": {"home/index.html"},
		"pages/list": {"pages/list.html"},
		"pages/detail": {"pages/detail.html"},
	}

	for name, tps := range templatesMap {
		s :=  append(layouts, tps...)
		t, err := template.ParseFS(templates.FS, s...)
		if err != nil {
			return nil, err
		}
		r.Add(name, t)
	}

	return r, nil
}