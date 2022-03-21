package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"goblog.com/pkg/jwtauth"
	"goblog.com/pkg/pagination"
	pb "goblog.com/service/article/proto/v1"
	"goblog.com/service/article/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type ArticleServiceServer struct {
	Repository *repository.Repository
	db *gorm.DB
	pb.UnimplementedArticleServiceServer
}

func NewArticleServiceServer(db *gorm.DB) pb.ArticleServiceServer {
	return &ArticleServiceServer{
		db: db,
		Repository: repository.NewRepository(db),
	}
}

func (s *ArticleServiceServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.Article, error) {
	//user, err := authUser(ctx, request.UserId)
	//if err != nil {
	//	return nil, err
	//}

	if request.PublishedTime == nil {
		request.PublishedTime = timestamppb.Now()
	}
	if request.UpdatedTime == nil {
		request.UpdatedTime = timestamppb.Now()
	}

	p := &repository.InsertParam{
		UserId:        0,
		Title:         request.Title,
		Content:       request.Content,
		Status:        int64(request.Status),
		PublishedTime: request.UpdatedTime.AsTime(),
		UpdatedTime:   request.UpdatedTime.AsTime(),
		Sort:          request.Sort,
	}

	id, err := s.Repository.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.UserPostId{
		Id: id,
	}, nil
}

func (s *ArticleServiceServer) Get(ctx context.Context, request *pb.ArticleId) (*pb.Article, error) {
	//user, err := authUser(ctx, request.UserId)
	//if err != nil {
	//	return nil, err
	//}

	i, err := s.Repository.Get(ctx, user.Id, request.Id)
	if err != nil {
		if err == repository.ErrNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return &pb.Article{
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


func (s *ArticleServiceServer) Update(ctx context.Context, request *pb.Article) (*pb.RowsAffected, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	p := &repository.UpdateParam{
		Title:         request.Title,
		Content:       request.Content,
		Status:        request.Status,
		PublishedTime: request.PublishedTime.AsTime(),
		UpdatedTime:   request.UpdatedTime.AsTime(),
		Sort:          request.Sort,
	}
	ra, err := s.Repository.Update(ctx, request.Id, user.Id, p)
	if err != nil {
		return nil, err
	}

	return &pb.Article{
		RowsAffected: ra,
	}, nil
}

func (s *ArticleServiceServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.Article, error) {
	user, err := authUser(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	ra, err := s.Repository.Delete(ctx, request.Id, user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *ArticleServiceServer) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	l := ctxzap.Extract(ctx)
	l.Info("222")
	grpc_ctxtags.Extract(ctx).Set("request", request)

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
