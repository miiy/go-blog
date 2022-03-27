package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/spf13/viper"
	"goblog.com/pkg/gin/middleware"
	"goblog.com/pkg/gin/template"
	"goblog.com/web/app/article"
	"goblog.com/web/app/book"
	"goblog.com/web/app/home"
	"goblog.com/web/resources/assets"
	"goblog.com/web/resources/templates"
	"net/http"
)

var v *viper.Viper

func RegisterRouter(r *gin.Engine, vv *viper.Viper) {
	v = vv
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

	home.Router(r)
	article.Router(r)
	book.Router(r)

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
	funMap := map[string]interface{}{
		"config": getConfig,
	}

	var tcs []template.Config
	for k, v := range templatesMap {
		c := template.Config{
			Name:    k,
			Files:   append(layouts, v...),
			FuncMap: funMap,
		}
		tcs = append(tcs, c)
	}

	tr, err := template.NewTemplateRender(templates.FS, tcs)
	if err != nil {
		panic(err)
	}
	return tr
}

func getConfig(key string) string {

	return v.GetString(key)
}