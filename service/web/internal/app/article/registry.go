package article

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	articlepb "goblog.com/api/article/v1"
)

type Article struct {
	router *gin.Engine
	logger *zap.Logger
	service *service
}

var article *Article

func NewArticle(router *gin.Engine, logger *zap.Logger, articleClient articlepb.ArticleServiceClient) *Article {
	article = &Article{
		router:  router,
		logger:  logger,
		service: NewService(articleClient, logger),
	}
	return article
}

var ProviderSet = wire.NewSet(NewArticle)
