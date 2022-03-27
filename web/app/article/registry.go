package article

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	//"goblog.com/pkg/application"
	"go.uber.org/zap"
)

type Article struct {
	router *gin.Engine
	logger *zap.Logger
}

var article *Article

func NewArticle(router *gin.Engine, logger *zap.Logger) *Article {
	article = &Article{
		router: router,
		logger: logger,
	}
	return article
}

var ProviderSet = wire.NewSet(NewArticle)
