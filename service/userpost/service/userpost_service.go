package service

import (
	"context"
	"database/sql"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "goblog.com/api/userpost/v1"
	"goblog.com/pkg/jwtauth"
	"goblog.com/pkg/pagination"
	"goblog.com/service/userpost/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UserPostServiceServer struct {
	Repository *repository.Repository
	db *sql.DB
	pb.UnimplementedUserPostServiceServer
}

func NewUserPostServiceServer(db *sql.DB) pb.UserPostServiceServer {
	return &UserPostServiceServer{
		db: db,
		Repository: repository.NewRepository(db),
	}
}

func authUser(ctx context.Context, requestUserId int64) (*jwtauth.AuthUser, error) {
	user, err := jwtauth.AuthUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if requestUserId != user.Id {
		return nil, status.New(codes.InvalidArgument, "invalid argument").Err()
	}
	return user, nil
}

func (s *UserPostServiceServer) CreateUserPost(ctx context.Context, request *pb.CreateUserPostRequest) (*pb.UserPostId, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	if request.PublishedTime == nil {
		request.PublishedTime = timestamppb.Now()
	}
	if request.UpdatedTime == nil {
		request.UpdatedTime = timestamppb.Now()
	}

	p := &repository.InsertUserPostParam{
		UserId:        user.Id,
		Title:         request.Title,
		Content:       request.Content,
		Status:        int64(request.Status),
		PublishedTime: request.UpdatedTime.AsTime(),
		UpdatedTime:   request.UpdatedTime.AsTime(),
		Sort:          request.Sort,
	}

	id, err := s.Repository.InsertUserPost(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.UserPostId{
		Id: id,
	}, nil
}

func (s *UserPostServiceServer) GetUserPost(ctx context.Context, request *pb.UserIdWithUserPostId) (*pb.UserPost, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	i, err := s.Repository.Get(ctx, user.Id, request.Id)
	if err != nil {
		if err == repository.ErrNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return &pb.UserPost{
		Id:            i.Id,
		UserId:        i.UserId,
		Title:         i.Title,
		Content:       i.Content,
		Status:        i.Status,
		PublishedTime: timestamppb.New(i.PublishedTime),
		UpdatedTime:   timestamppb.New(i.UpdatedTime),
		Sort:          i.Sort,
		CreatedAt:     timestamppb.New(i.CreatedAt),
		UpdatedAt:     timestamppb.New(i.UpdatedAt),
	}, nil
}


func (s *UserPostServiceServer) UpdateUserPost(ctx context.Context, request *pb.UpdateUserPostRequest) (*pb.RowsAffected, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	p := &repository.UpdateUserPostParam{
		Title:         request.Title,
		Content:       request.Content,
		Status:        request.Status,
		PublishedTime: request.PublishedTime.AsTime(),
		UpdatedTime:   request.UpdatedTime.AsTime(),
		Sort:          request.Sort,
	}
	ra, err := s.Repository.UpdateUserPost(ctx, request.Id, user.Id, p)
	if err != nil {
		return nil, err
	}

	return &pb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *UserPostServiceServer) UpdateUserPostSort(ctx context.Context, request *pb.UpdateUserPostSortRequest) (*pb.RowsAffected, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	var sort int64
	if request.SortType == pb.UpdateUserPostSortRequest_TOP {
		sort = time.Now().Unix()
	}
	if request.SortType == pb.UpdateUserPostSortRequest_BOTTOM {
		sort = 0
	}
	if request.SortType == pb.UpdateUserPostSortRequest_BOTTOM {
		sort = request.Sort
	}

	ra, err := s.Repository.UpdateSort(ctx, request.Id, user.Id, sort)
	if err != nil {
		return nil, err
	}

	return &pb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *UserPostServiceServer) DeleteUserPost(ctx context.Context, request *pb.UserIdWithUserPostId) (*pb.RowsAffected, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	ra, err := s.Repository.DeleteUserPost(ctx, request.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *UserPostServiceServer) ListUserPost(ctx context.Context, request *pb.ListUserPostRequest) (*pb.ListUserPostResponse, error) {
	l := ctxzap.Extract(ctx)
	l.Info("222")
	grpcctxtags.Extract(ctx).Set("request", request)

	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	// validate

	// count
	total, err := s.Repository.ListCount(ctx, user.Id, request.Keyword)
	if err != nil {
		return nil, err
	}

	// pagination
	pg := pagination.NewPagination(request.Page, request.PerPage, total)

	// list
	tags, err := s.Repository.List(ctx, user.Id, pg.PerPage, pg.From, request.Keyword)
	if err != nil {
		return nil, err
	}
	var items []*pb.UserPost
	for _, v  := range tags {
		item := &pb.UserPost{
			Id:            v.Id,
			UserId:        v.UserId,
			Title:         v.Title,
			Content:       v.Content,
			Status:        v.Status,
			PublishedTime: timestamppb.New(v.PublishedTime),
			UpdatedTime:   timestamppb.New(v.UpdatedTime),
			Sort:          v.Sort,
			CreatedAt:     timestamppb.New(v.CreatedAt),
			UpdatedAt:     timestamppb.New(v.UpdatedAt),
		}
		items = append(items, item)
	}

	return &pb.ListUserPostResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Data:        items,
	}, nil
}
