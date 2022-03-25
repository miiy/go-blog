package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"goblog.com/pkg/gin/middleware"
	"goblog.com/web/pkg/config"
	"goblog.com/web/pkg/template"
	"goblog.com/web/resources/assets"
	"net/http"
)

func RegisterRouter(r *gin.Engine, appOpts config.AppOptions) {
	// template
	r.HTMLRender = htmlRender()

	// assets
	r.StaticFS("/assets", http.FS(assets.FS))

	// favicon
	faviconHandler := func(c *gin.Context) {
		c.FileFromFS("favicon.ico", http.FS(assets.FS))
	}
	r.HEAD("/favicon.ico", faviconHandler)
	r.GET("/favicon.ico", faviconHandler)

	// uploads
	r.Static("/uploads", "./storage/uploads")


	// index
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home/index", gin.H{
			"PageTitle": "Home",
			"Content":   "Hello, world.",
			"Header":    "header.",
			"App": appOpts,
		})
	})

	// article
	r.GET("/articles", func(c *gin.Context) {
		c.HTML(200, "article/list", gin.H{
			"PageTitle": "Article list",
		})
	})
	r.GET("/articles/:id", func(c *gin.Context) {
		c.HTML(200, "article/detail", gin.H{
			"PageTitle": "Article detail",
		})
	})
	// book
	r.GET("/books", func(c *gin.Context) {
		c.HTML(200, "book/list", gin.H{
			"PageTitle": "Book list",
		})
	})
	// book detail
	r.GET("/books/:id", func(c *gin.Context) {
		c.HTML(200, "book/detail", gin.H{
			"PageTitle": "Book detail",
		})
	})

	// pages
	r.GET("/pages/list", func(c *gin.Context) {
		c.HTML(200, "pages/list", gin.H{})
	})
	r.GET("/pages/detail", func(c *gin.Context) {
		c.HTML(200, "pages/detail", gin.H{})
	})

	// middleware
	r.Use(middleware.RequestInfo())

}

func htmlRender() render.HTMLRender {
	layouts := []string{"layout/layout.html", "layout/header.html", "layout/footer.html"}
	templatesMap := map[string][]string{
		"home/index":     {"home/index.html"},
		"article/detail": {"article/detail.html"},
		"article/list":   {"article/list.html"},
		"book/detail":    {"book/detail.html"},
		"book/list":      {"book/list.html"},
		"pages/list":     {"pages/list.html"},
		"pages/detail":   {"pages/detail.html"},
	}
	for k, v := range templatesMap {
		templatesMap[k] = append(layouts, v...)
	}
	tr, err := template.NewTemplateRender(templatesMap)
	if err != nil {
		panic(err)
	}
	return tr
}