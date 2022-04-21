package service

import (
	"context"
	"go.uber.org/zap"
	articlepb "goblog.com/api/article/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/article/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type ArticleServiceServer struct {
	Repository *repository.ArticleRepository
	db *gorm.DB
	articlepb.UnimplementedArticleServiceServer
}

func NewArticleServiceServer(db *gorm.DB, logger *zap.Logger) articlepb.ArticleServiceServer {
	return &ArticleServiceServer{
		db: db,
		Repository: repository.NewArticleRepository(db, logger),
	}
}

func (s *ArticleServiceServer) CreateArticle(ctx context.Context, request *articlepb.CreateArticleRequest) (*articlepb.Article, error) {
	article := request.GetArticle()
	if article.PublishedTime == nil {
		article.PublishedTime = timestamppb.Now()
	}
	if article.UpdatedTime == nil {
		article.UpdatedTime = timestamppb.Now()
	}

	a := &repository.Article{
		UserId:          0,
		CategoryId:      article.CategoryId,
		Title:           article.Title,
		MetaTitle:       article.MetaTitle,
		MetaDescription: article.MetaDescription,
		PublishedTime:   article.PublishedTime.AsTime(),
		UpdatedTime:     article.UpdatedTime.AsTime(),
		FromText:        article.FromText,
		FromUrl:         article.FromUrl,
		Summary:         article.Summary,
		Content:         article.Content,
		Status:          int(article.Status),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	_, err := s.Repository.Create(ctx, a)
	if err != nil {
		return nil, err
	}

	return articleToProto(a), nil
}

func (s *ArticleServiceServer) BatchCreateArticles(ctx context.Context, request *articlepb.BatchCreateArticlesRequest) (*articlepb.BatchCreateArticlesResponse, error) {
	return nil, nil
}

func (s *ArticleServiceServer) GetArticle(ctx context.Context, request *articlepb.GetArticleRequest) (*articlepb.Article, error) {
	article, err := s.Repository.First(ctx, request.Id, "*")
	if err != nil {
		if err == repository.ErrRecordNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return articleToProto(article), nil
}


func (s *ArticleServiceServer) UpdateArticle(ctx context.Context, request *articlepb.UpdateArticleRequest) (*articlepb.Article, error) {
	article := request.GetArticle()

	a := &repository.Article{
		UserId:          0,
		CategoryId:      article.CategoryId,
		Title:           article.Title,
		MetaTitle:       article.MetaTitle,
		MetaDescription: article.MetaDescription,
		PublishedTime:   article.PublishedTime.AsTime(),
		UpdatedTime:     article.UpdatedTime.AsTime(),
		FromText:        article.FromText,
		FromUrl:         article.FromUrl,
		Summary:         article.Summary,
		Content:         article.Content,
		Status:          int(article.Status),
	}
	ra, err := s.Repository.Update(ctx, request.Id, a)
	if err != nil {
		return nil, err
	}

	return articleToProto(ra), nil
}

func (s *ArticleServiceServer) DeleteArticle(ctx context.Context, request *articlepb.DeleteArticleRequest) (*emptypb.Empty, error) {
	err := s.Repository.Delete(ctx, request.Id)
	return nil, err
}

func (s *ArticleServiceServer) ListArticles(ctx context.Context, request *articlepb.ListArticlesRequest) (*articlepb.ListArticlesResponse, error) {
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
	var items []*articlepb.Article
	for _, v  := range articles {
		items = append(items, articleToProto(v))
	}

	return &articlepb.ListArticlesResponse{
		Total:       pg.Total,
		PageSize:    pg.PerPage,
		CurrentPage: pg.CurrentPage,
		Articles:    items,
	}, nil
}

func articleToProto(a *repository.Article) *articlepb.Article {
	article := &articlepb.Article{
		Id:              a.Id,
		UserId:          a.UserId,
		Title:           a.Title,
		MetaTitle:       a.MetaTitle,
		MetaDescription: a.MetaDescription,
		PublishedTime:   timestamppb.New(a.PublishedTime),
		UpdatedTime:     timestamppb.New(a.UpdatedTime),
		FromText:        a.FromText,
		FromUrl:         a.FromUrl,
		Summary:         a.Summary,
		Content:         a.Content,
		Status:          articlepb.Article_ArticleStatus(a.Status),
		CreateTime:      timestamppb.New(a.CreatedAt),
		UpdateTime:      timestamppb.New(a.UpdatedAt),
	}

	return article
}