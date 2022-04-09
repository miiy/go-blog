package service

import (
	"context"
	userpostpb "goblog.com/api/userpost/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/userpost/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserPostServiceServer) CreateUserTag(ctx context.Context, request *userpostpb.CreateUserTagRequest) (*userpostpb.UserTagId, error) {

	p := &repository.InsertUserTagParam{
		UserId:        request.UserId,
		Name:          request.Name,
	}

	id, err := s.Repository.InsertUserTag(ctx, p)
	if err != nil {
		return nil, err
	}
	return &userpostpb.UserTagId{
		Id: id,
	}, nil
}

func (s *UserPostServiceServer) GetUserTag(ctx context.Context, request *userpostpb.UserIdWithUserTagId) (*userpostpb.UserTag, error) {
	i, err := s.Repository.Get(ctx, request.UserId, request.Id)
	if err != nil {
		if err == repository.ErrNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return &userpostpb.UserTag{
		Id:            i.Id,
		UserId:        i.UserId,
		Name:          i.Name,
		CreatedAt:     timestamppb.New(i.CreatedAt),
		UpdatedAt:     timestamppb.New(i.UpdatedAt),
	}, nil
}


func (s *UserPostServiceServer) UpdateUserTag(ctx context.Context, request *userpostpb.UpdateUserTagRequest) (*userpostpb.RowsAffected, error) {
	return &userpostpb.RowsAffected{
		RowsAffected: 0,
	}, nil
}

func (s *UserPostServiceServer) DeleteUserTag(ctx context.Context, request *userpostpb.UserIdWithUserTagId) (*userpostpb.RowsAffected, error) {
	return &userpostpb.RowsAffected{
		RowsAffected: 0,
	}, nil
}

func (s *UserPostServiceServer) ListUserTag(ctx context.Context, request *userpostpb.ListUserTagRequest) (*userpostpb.ListUserTagResponse, error) {
	// validate

	userId := request.UserId
	// count
	total, err := s.Repository.ListCount(ctx, userId)
	if err != nil {
		return nil, err
	}

	// pagination
	pg := pagination.NewPagination(request.Page, request.PerPage, total)

	// list
	tags, err := s.Repository.List(ctx, userId, pg.PerPage, pg.From)
	if err != nil {
		return nil, err
	}
	var items []*userpostpb.UserTag
	for _, v  := range tags {
		item := &userpostpb.UserTag{
			Id:            v.Id,
			UserId:        v.UserId,
			Name:          v.Name,
			CreatedAt:     timestamppb.New(v.CreatedAt),
			UpdatedAt:     timestamppb.New(v.UpdatedAt),
		}
		items = append(items, item)
	}

	return &userpostpb.ListUserTagResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Data:        items,
	}, nil
}