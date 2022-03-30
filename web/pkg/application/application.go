package application

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/web/app/article"
	"goblog.com/web/app/book"
	"goblog.com/web/pkg/config"
)

type Application struct {
	Config  *config.Config
	Router  *gin.Engine
	Logger  *zap.Logger
	Article *article.Article
	Book    *book.Book
}

func NewApplication(config *config.Config, router *gin.Engine, logger *zap.Logger, article *article.Article, book *book.Book) *Application {
	return &Application{
		Config: config,
		Logger: logger,
		Router: router,
		Article: article,
		Book: book,
	}
}

var ProviderSet = wire.NewSet(NewApplication)

func (app *Application) RegisterRouter(funcs ...func(r *gin.Engine)) {
	for _, f := range funcs {
		f(app.Router)
	}
}