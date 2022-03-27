// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"goblog.com/pkg/gin"
	"goblog.com/pkg/logger"
	"goblog.com/web/app/article"
	"goblog.com/web/app/book"
	"goblog.com/web/pkg/application"
)

// Injectors from wre.go:

func InitApplication() (*application.Application, func(), error) {
	options := gin.NewOptions()
	engine, err := gin.NewGin(options)
	if err != nil {
		return nil, nil, err
	}
	loggerOptions := logger.NewOptions()
	zapLogger, cleanup, err := logger.NewLogger(loggerOptions)
	if err != nil {
		return nil, nil, err
	}
	articleArticle := article.NewArticle(engine, zapLogger)
	bookBook := book.NewBook(engine, zapLogger)
	applicationApplication := application.NewApplication(engine, zapLogger, articleArticle, bookBook)
	return applicationApplication, func() {
		cleanup()
	}, nil
}

// wre.go:

var providerSet = wire.NewSet(application.ProviderSet, gin.ProviderSet, logger.ProviderSet, article.ProviderSet, book.ProviderSet)
