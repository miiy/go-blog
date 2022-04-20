package server

import (
	"context"
	"database/sql"
	feedbackpb "goblog.com/api/feedback/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/feedback/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FeedbackServer struct {
	Repository *repository.Repository
	feedbackpb.UnimplementedFeedbackServiceServer
}

func NewFeedbackServer(db *sql.DB) feedbackpb.FeedbackServiceServer {
	r := repository.NewRepository(db)
	return &FeedbackServer{
		Repository: r,
	}
}

func (s *FeedbackServer) CreateFeedback(ctx context.Context, request *feedbackpb.CreateFeedbackRequest) (*feedbackpb.Feedback, error) {
	// validate
	if request.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "content can not empty")
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

func (s *FeedbackServer) DeleteFeedback(ctx context.Context, request *feedbackpb.DeleteFeedbackRequest) (*emptypb.Empty, error) {
	err := s.Repository.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *FeedbackServer) ListFeedbacks(ctx context.Context, request *feedbackpb.ListFeedbacksRequest) (*feedbackpb.ListFeedbacksResponse, error) {
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