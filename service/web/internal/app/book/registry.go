package book

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	bookpb "goblog.com/api/book/v1"
)

type Book struct {
	router *gin.Engine
	logger *zap.Logger
	service *service
}

var book *Book

func NewArticle(router *gin.Engine, logger *zap.Logger, bookClient bookpb.BookServiceClient) *Book {
	book = &Book{
		router:  router,
		logger:  logger,
		service: NewService(bookClient, logger),
	}
	return book
}

var ProviderSet = wire.NewSet(NewArticle)
