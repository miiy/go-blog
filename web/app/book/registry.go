package book

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Book struct {
	router *gin.Engine
	logger *zap.Logger
}

var book *Book

func NewBook(router *gin.Engine, logger *zap.Logger) *Book {
	book = &Book{
		router: router,
		logger: logger,
	}
	return book
}

var ProviderSet = wire.NewSet(NewBook)
