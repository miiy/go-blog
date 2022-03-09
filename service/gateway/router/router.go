package router

import (
	"github.com/gin-gonic/gin"
	"github.com/miiy/go-blog/pkg/gin/middleware"
)

func registerRouter(r *gin.Engine) {
	// template
	r.LoadHTMLGlob("internal/app/**/templates/*.tmpl")

	// favicon
	r.StaticFile("/favicon.ico", "./web/favicon.ico")
	// statics
	r.Static("/statics", "./web/statics")
	// uploads
	r.Static("/uploads", "./storage/uploads")

	// middleware
	r.Use(middleware.RequestInfo())

}