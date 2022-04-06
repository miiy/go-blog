package article

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.uber.org/zap"
	//"goblog.com/pkg/application"
	articleClient "goblog.com/service/article/proto/v1"
)

type Article struct {
	router *gin.Engine
	logger *zap.Logger
	service *service
}

var article *Article

func NewArticle(router *gin.Engine, logger *zap.Logger) *Article {

	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("dit not connect: %v", zap.Error(err))
	}
	//defer conn.Close()

	ac := articleClient.NewArticleServiceClient(conn)

	article = &Article{
		router: router,
		logger: logger,
		service: NewService(ac, logger),
	}
	return article
}

var ProviderSet = wire.NewSet(NewArticle)
