// +build wireinject

package main

import (
	"github.com/google/wire"
	"goblog.com/pkg/gin"
	"goblog.com/pkg/logger"
	"goblog.com/web/app/article"
	"goblog.com/web/app/book"
	"goblog.com/web/pkg/application"
)

var providerSet = wire.NewSet(
	application.ProviderSet,
	gin.ProviderSet,
	logger.ProviderSet,
	article.ProviderSet,
	book.ProviderSet,
)

func InitApplication() (*application.Application, func(), error) {
	panic(wire.Build(providerSet))
}