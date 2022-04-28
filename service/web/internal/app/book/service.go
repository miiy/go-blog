package book

import (
	"context"
	"go.uber.org/zap"
	bookpb "goblog.com/api/book/v1"
	"time"
)

type service struct {
	 client bookpb.BookServiceClient
	 logger *zap.Logger
}

func NewService(client bookpb.BookServiceClient, logger *zap.Logger) *service {
	return &service{
		client: client,
		logger: logger,
	}
}

func (s *service) ListBooks(cid, page, pageSize int) (*bookpb.ListBooksResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := bookpb.ListBooksRequest{
		CategoryId: int64(cid),
		Page: int64(page),
		PageSize: int64(pageSize),
	}
	resp, err := s.client.ListBooks(ctx, &req)
	if err != nil {
		s.logger.Error("client.ListBooks(_) = _, %v", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *service) GetBook(id int) (*bookpb.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := bookpb.GetBookRequest{
		Id: int64(id),
	}
	resp, err := s.client.GetBook(ctx, &req)
	if err != nil {
		s.logger.Error("client.GetBook(_) = _, %v", zap.Error(err))
		return nil, err
	}

	return resp, nil
}
