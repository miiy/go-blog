// +build wireinject

package main

import (
	"github.com/google/wire"
	"goblog.com/pkg/environment"
	"goblog.com/pkg/gin"
	"goblog.com/pkg/logger"
	"goblog.com/service/web/app/article"
	"goblog.com/service/web/app/book"
	"goblog.com/service/web/pkg/application"
	"goblog.com/service/web/pkg/config"
)

func InitApplication(conf string) (*application.Application, func(), error) {
	panic(
		wire.Build(
			application.ProviderSet,
			config.ProviderSet,
			providerGinOption,
			gin.ProviderSet,
			logger.ProviderSet,
			providerLoggerOption,
			article.ProviderSet,
			book.ProviderSet,
		),
	)
}

func providerGinOption() []gin.Option {
	return []gin.Option{
		gin.WithEnv(environment.DEVELOPMENT),
	}
}

func providerLoggerOption() []logger.Option {
	return []logger.Option{
		logger.WithEnv(environment.DEVELOPMENT),
	}
}