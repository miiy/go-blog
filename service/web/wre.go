// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	articlepb "goblog.com/api/article/v1"
	"goblog.com/pkg/environment"
	"goblog.com/pkg/gin"
	"goblog.com/pkg/logger"
	"goblog.com/service/web/app/article"
	"goblog.com/service/web/app/book"
	"goblog.com/service/web/pkg/application"
	"goblog.com/service/web/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
			providerArticleClient,
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

func providerArticleClient(logger *zap.Logger) (articlepb.ArticleServiceClient, func()) {
	// Set up a connection to the server.

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("dit not connect: %v", zap.Error(err))
	}
	// defer conn.Close()

	ac := articlepb.NewArticleServiceClient(conn)
	return ac, func() {
		defer  conn.Close()
	}
}