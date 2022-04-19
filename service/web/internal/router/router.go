package router

import (
	"github.com/gin-gonic/gin"
	"goblog.com/pkg/gin/middleware"
	"goblog.com/pkg/gin/template"
	"goblog.com/service/web/internal/app/article"
	"goblog.com/service/web/internal/app/book"
	"goblog.com/service/web/internal/app/home"
	"goblog.com/service/web/internal/pkg/config"
	"goblog.com/service/web/internal/resources/assets"
	"goblog.com/service/web/internal/resources/templates"
	"net/http"
)

func Router(r *gin.Engine) {
	setDefaultRouter(r)

	// template
	t := template.NewTemplate()
	t.AddFunc("config", func(key string) string {
		return config.GetString(key)
	})
	t.AddTemplate(templates.FS, home.Templates())
	t.AddTemplate(templates.FS, article.Templates())
	t.AddTemplate(templates.FS, book.Templates())

	r.HTMLRender = t.Render

	// modules router
	home.Router(r)
	article.Router(r)
	book.Router(r)

	// middleware
	r.Use(middleware.RequestInfo())
}

func setDefaultRouter(r *gin.Engine)  {
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
}
