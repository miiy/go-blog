package article

import (
	"context"
	"go.uber.org/zap"
	articlepb "goblog.com/api/article/v1"
	"time"
)

type service struct {
	 client articlepb.ArticleServiceClient
	 logger *zap.Logger
}

func NewService(client articlepb.ArticleServiceClient, logger *zap.Logger) *service {
	return &service{
		client: client,
		logger: logger,
	}
}

func (s *service) ArticleList() (*articlepb.ListArticlesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := articlepb.ListArticlesRequest{
		Page: 1,
		PageSize: 10,
	}
	resp, err := s.client.ListArticles(ctx, &req)
	if err != nil {
		s.logger.Error("client.GetArticle(_) = _, %v", zap.Error(err))
		return nil, err
	}

	return resp, nil
}