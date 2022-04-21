// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	articlepb "goblog.com/api/article/v1"
	bookpb "goblog.com/api/book/v1"
	"goblog.com/pkg/environment"
	"goblog.com/pkg/gin"
	"goblog.com/pkg/logger"
	"goblog.com/service/web/internal/app/article"
	"goblog.com/service/web/internal/app/book"
	"goblog.com/service/web/internal/pkg/application"
	"goblog.com/service/web/internal/pkg/config"
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
			providerBookClient,
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
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("dit not connect: %v", zap.Error(err))
	}

	c := articlepb.NewArticleServiceClient(conn)
	return c, func() {
		defer  conn.Close()
	}
}

func providerBookClient(logger *zap.Logger) (bookpb.BookServiceClient, func()) {
	conn, err := grpc.Dial("127.0.0.1:50053", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("dit not connect: %v", zap.Error(err))
	}

	c := bookpb.NewBookServiceClient(conn)
	return c, func() {
		defer  conn.Close()
	}
}