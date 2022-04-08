package service

import (
	"context"
	"database/sql"
	pb "goblog.com/service/usertag/proto"
	"goblog.com/service/usertag/repository"
	"goblog.com/pkg/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserTagServiceServer struct {
	Repository *repository.Repository
	db *sql.DB
	pb.UnimplementedUserTagServiceServer
}

func NewUserTagServiceServer(db *sql.DB) pb.UserTagServiceServer {
	repository := repository.NewRepository(db)
	return &UserTagServiceServer{
		db: db,
		Repository: repository,
	}
}

func (s *UserTagServiceServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.UserTagId, error) {

	p := &repository.InsertParam{
		UserId:        request.UserId,
		Name:          request.Name,
	}

	id, err := s.Repository.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.UserTagId{
		Id: id,
	}, nil
}

func (s *UserTagServiceServer) Get(ctx context.Context, request *pb.UserIdWithUserTagId) (*pb.UserTag, error) {
	i, err := s.Repository.Get(ctx, request.UserId, request.Id)
	if err != nil {
		if err == repository.ErrNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return &pb.UserTag{
		Id:            i.Id,
		UserId:        i.UserId,
		Name:          i.Name,
		CreatedAt:     timestamppb.New(i.CreatedAt),
		UpdatedAt:     timestamppb.New(i.UpdatedAt),
	}, nil
}


func (s *UserTagServiceServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.RowsAffected, error) {
	return &pb.RowsAffected{
		RowsAffected: 0,
	}, nil
}

func (s *UserTagServiceServer) Delete(ctx context.Context, request *pb.UserIdWithUserTagId) (*pb.RowsAffected, error) {
	return &pb.RowsAffected{
		RowsAffected: 0,
	}, nil
}

func (s *UserTagServiceServer) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
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
	var items []*pb.UserTag
	for _, v  := range tags {
		item := &pb.UserTag{
			Id:            v.Id,
			UserId:        v.UserId,
			Name:          v.Name,
			CreatedAt:     timestamppb.New(v.CreatedAt),
			UpdatedAt:     timestamppb.New(v.UpdatedAt),
		}
		items = append(items, item)
	}

	return &pb.ListResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Data:        items,
	}, nil
}