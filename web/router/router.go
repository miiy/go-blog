package router

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"goblog.com/pkg/gin/middleware"
	"goblog.com/web/pkg/config"
	"goblog.com/web/resources/assets"
	"goblog.com/web/resources/templates"
	"html/template"
	"net/http"
	"path"
)

func RegisterRouter(r *gin.Engine, appOpts config.AppOptions) {


	// template
	tr, err := createTemplateRender()
	if err != nil {
		panic(err)
	}
	r.HTMLRender = tr

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
func unescaped (x string) interface{} { return template.HTML(x) }

func createTemplateRender() (multitemplate.Renderer, error) {
	r := multitemplate.NewRenderer()

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

	for name, tps := range templatesMap {
		// template name see template.ParseFiles
		tName := path.Base(layouts[0])
		t, err := template.New(tName).Funcs(template.FuncMap{
			"unescaped": unescaped,
		}).ParseFS(templates.FS, append(layouts, tps...)...)
		fmt.Println(t.Templates())
		if err != nil {
			return nil, err
		}
		r.Add(name, t)
	}

	return r, nil
}
