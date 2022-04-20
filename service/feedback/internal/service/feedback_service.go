package service

import (
	"context"
	"database/sql"
	"errors"
	feedbackpb "goblog.com/api/feedback/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/feedback/internal/repository"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedbackServiceServer struct {
	Repository *repository.Repository
	feedbackpb.UnimplementedFeedbackServiceServer
}

func NewFeedbackServiceServer(db *sql.DB) feedbackpb.FeedbackServiceServer {
	r := repository.NewRepository(db)
	return &FeedbackServiceServer{
		Repository: r,
	}
}

func (s *FeedbackServiceServer) Create(ctx context.Context, request *feedbackpb.CreateFeedbackRequest) (*feedbackpb.Feedback, error) {
	// validate
	if request.Content == "" {
		return nil, errors.New("content can not empty")
	}

	p := &repository.InsertParam{
		UserId:  request.UserId,
		Content: request.Content,
	}

	feedback, err := s.Repository.Insert(ctx, p)
	if err != nil {
		return nil, err
	}

	return feedbackToProto(feedback), nil
}

func (s *FeedbackServiceServer) Delete(ctx context.Context, request *feedbackpb.DeleteFeedbackRequest) (*emptypb.Empty, error) {
	err := s.Repository.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *FeedbackServiceServer) List(ctx context.Context, request *feedbackpb.ListFeedbacksRequest) (*feedbackpb.ListFeedbacksResponse, error) {
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
		item := feedbackToProto(v)
		items = append(items, item)
	}

	return &feedbackpb.ListFeedbacksResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Feedbacks:   items,
	}, nil
}

func feedbackToProto(feedback *repository.Feedback) *feedbackpb.Feedback {
	return &feedbackpb.Feedback{
		Id:        feedback.Id,
		UserId:    feedback.UserId,
		Content:   feedback.Content,
		CreatedAt: timestamppb.New(feedback.CreatedAt),
		UpdatedAt: timestamppb.New(feedback.UpdatedAt),
		DeletedAt: timestamppb.New(feedback.DeletedAt.Time),
	}
}