package service

import (
	"context"
	"database/sql"
	feedbackpb "goblog.com/api/feedback/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/feedback/internal/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedbackServiceServer struct {
	Repository *repository.Repository
	feedbackpb.UnimplementedFeedbackServiceServer
}

func NewFeedbackServiceServer(db *sql.DB) feedbackpb.FeedbackServiceServer {
	repository := repository.NewRepository(db)
	return &FeedbackServiceServer{
		Repository: repository,
	}
}

func (s *FeedbackServiceServer) Create(ctx context.Context, request *feedbackpb.CreateFeedback) (*feedbackpb.FeedbackId, error) {
	// validate
	p := &repository.InsertParam{
		UserId:  request.UserId,
		Content: request.Content,
	}
	id, err := s.Repository.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	return &feedbackpb.FeedbackId{
		Id: id,
	}, nil
}

func (s *FeedbackServiceServer) Delete(ctx context.Context, request *feedbackpb.FeedbackId) (*feedbackpb.RowsAffected, error) {
	ra, err := s.Repository.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &feedbackpb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *FeedbackServiceServer) List(ctx context.Context, request *feedbackpb.ListRequest) (*feedbackpb.ListResponse, error) {
	// validate

	// count
	total, err := s.Repository.ListCount(ctx)
	if err != nil {
		return nil, err
	}

	// pagination
	pg := pagination.NewPagination(request.Page, request.PerPage, total)

	// list
	tags, err := s.Repository.List(ctx, request.UserId, pg.PerPage, pg.From)
	if err != nil {
		return nil, err
	}
	var items []*feedbackpb.Feedback
	for _, v  := range tags {
		item := &feedbackpb.Feedback{
			Id:        v.Id,
			UserId:    v.UserId,
			Content:   v.Content,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		}
		items = append(items, item)
	}

	return &feedbackpb.ListResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Data:        items,
	}, nil
}
