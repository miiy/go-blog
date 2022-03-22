package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
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

func NewArticleServiceServer(db *gorm.DB, logger *zap.Logger) pb.ArticleServiceServer {
	return &ArticleServiceServer{
		db: db,
		Repository: repository.NewRepository(db, logger),
	}
}

func (s *ArticleServiceServer) Create(ctx context.Context, request *pb.Article) (*pb.Article, error) {
	//user, err := authUser(ctx, request.UserId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if request.PublishedTime == nil {
	//	request.PublishedTime = timestamppb.Now()
	//}
	//if request.UpdatedTime == nil {
	//	request.UpdatedTime = timestamppb.Now()
	//}

	a := &repository.Article{
		UserId:          0,
		CategoryId:      request.CategoryId,
		Title:           request.Title,
		MetaTitle:       request.MetaTitle,
		MetaDescription: request.MetaDescription,
		PublishedTime:   request.PublishedTime.AsTime(),
		UpdatedTime:     request.UpdatedTime.AsTime(),
		FromTest:        request.FromText,
		FromUrl:         request.FromUrl,
		Summary:         request.Summary,
		Content:         request.Content,
		Status:          int(request.Status),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	_, err := s.Repository.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	articlePb, err := articleToProto(a)
	if err != nil {
		return nil, err
	}
	return articlePb, nil
}

func (s *ArticleServiceServer) Get(ctx context.Context, request *pb.GetArticleRequest) (*pb.Article, error) {
	a, err := s.Repository.First(ctx, request.Id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	articlePb, err := articleToProto(a)
	if err != nil {
		return nil, err
	}
	return articlePb, nil
}


func (s *ArticleServiceServer) Update(ctx context.Context, request *pb.Article) (*pb.Article, error) {
	a := &repository.Article{
		UserId:          0,
		CategoryId:      request.CategoryId,
		Title:           request.Title,
		MetaTitle:       request.MetaTitle,
		MetaDescription: request.MetaDescription,
		PublishedTime:   request.PublishedTime.AsTime(),
		UpdatedTime:     request.UpdatedTime.AsTime(),
		FromTest:        request.FromText,
		FromUrl:         request.FromUrl,
		Summary:         request.Summary,
		Content:         request.Content,
		Status:          int(request.Status),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	ra, err := s.Repository.Update(ctx, request.Id, a)
	if err != nil {
		return nil, err
	}

	articlePb, err := articleToProto(ra)
	if err != nil {
		return nil, err
	}
	return articlePb, nil
}

func (s *ArticleServiceServer) Delete(ctx context.Context, request *pb.DeleteArticleRequest) error {
	err := s.Repository.Delete(ctx, request.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleServiceServer) List(ctx context.Context, request *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
	l := ctxzap.Extract(ctx)
	l.Info("222")
	grpc_ctxtags.Extract(ctx).Set("request", request)

	//user, err := authUser(ctx, request.UserId)
	//if err != nil {
	//	return nil, err
	//}
	// validate

	// count
	total, err := s.Repository.FindCount(ctx)
	if err != nil {
		return nil, err
	}

	// pagination
	pg := pagination.NewPagination(request.Page, request.PageSize, total)

	// list
	articles, err := s.Repository.Find(ctx, pg.PerPage, pg.From)
	if err != nil {
		return nil, err
	}
	var items []*pb.Article
	for _, v  := range articles {
		item, err := articleToProto(v)
		if err != nil {
			continue
		}
		items = append(items, item)
	}

	return &pb.ListArticlesResponse{
		Total:       pg.Total,
		PageSize:    pg.PerPage,
		CurrentPage: pg.CurrentPage,
		Articles:    items,
	}, nil
}

func articleToProto(a *repository.Article) (*pb.Article, error) {
	article := &pb.Article{
		Id:              a.Id,
		UserId:          a.UserId,
		Title:           a.Title,
		MetaTitle:       a.MetaTitle,
		MetaDescription: a.MetaDescription,
		PublishedTime:   timestamppb.New(a.PublishedTime),
		UpdatedTime:     timestamppb.New(a.UpdatedTime),
		FromText:        a.FromTest,
		FromUrl:         a.FromUrl,
		Summary:         a.Summary,
		Content:         a.Content,
		Status:          pb.Article_ArticleStatus(a.Status),
		CreateTime:      timestamppb.New(a.CreatedAt),
		UpdateTime:      timestamppb.New(a.UpdatedAt),
	}

	return article, nil
}