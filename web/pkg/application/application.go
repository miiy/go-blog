package application

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/web/app/article"
	"goblog.com/web/app/book"
)

type Application struct {
	Env string
	Locale string
	Debug bool

	Version string

	Router *gin.Engine
	Logger *zap.Logger
	Article *article.Article
	Book    *book.Book
}

func NewApplication(router *gin.Engine, logger *zap.Logger, article *article.Article, book *book.Book) *Application {
	return &Application{
		Env: "",
		Locale: "",
		Debug: false,

		Version: "",

		Logger: logger,
		Router: router,
		Article: article,
		Book: book,
	}
}

var ProviderSet = wire.NewSet(NewApplication)