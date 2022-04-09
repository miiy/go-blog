package v1

import (
	"context"
	"database/sql"
	tagpb "goblog.com/api/tag/v1"
	"goblog.com/pkg/pagination"
	"goblog.com/service/tag/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Options struct {
	Debug bool
}

type TagServiceServer struct {
	Options *Options
	Repository *repository.Repository
	tagpb.UnimplementedTagServiceServer
}

func NewTagServiceServer(o *Options, db *sql.DB) tagpb.TagServiceServer {
	repository := repository.NewRepository(db)
	return &TagServiceServer{
		Options: o,
		Repository: repository,
	}
}

func (s *TagServiceServer) Create(ctx context.Context, request *tagpb.CreateTag) (*tagpb.TagId, error) {
	// validate
	p := &repository.InsertParam{
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
	}
	id, err := s.Repository.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	return &tagpb.TagId{
		Id: id,
	}, nil
}

func (s *TagServiceServer) Get(ctx context.Context, request *tagpb.TagId) (*tagpb.Tag, error) {
	t, err := s.Repository.Get(ctx, request.Id)
	if err != nil {
		if err == repository.ErrorNotFound {
			st := status.New(codes.NotFound, "tag was not found")
			return nil, st.Err()
		}
		st := status.New(codes.Internal, "internal server error")
		return nil, st.Err()
	}

	return &tagpb.Tag{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
	}, nil
}

func (s *TagServiceServer) Update(ctx context.Context, request *tagpb.UpdateTag) (*tagpb.RowsAffected, error) {
	p := &repository.UpdateParam{
		Name:        request.Name,
		Description: request.Description,
		Status:      request.Status,
	}
	ra, err := s.Repository.Update(ctx, request.Id, p)
	if err != nil {
		return nil, err
	}
	return &tagpb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *TagServiceServer) Delete(ctx context.Context, request *tagpb.TagId) (*tagpb.RowsAffected, error) {
	ra, err := s.Repository.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	return &tagpb.RowsAffected{
		RowsAffected: ra,
	}, nil
}

func (s *TagServiceServer) List(ctx context.Context, request *tagpb.ListRequest) (*tagpb.ListResponse, error) {
	// validate

	// count
	total, err := s.Repository.ListCount(ctx, request.Keyword)
	if err != nil {
		return nil, err
	}

	// pagination
	pg := pagination.NewPagination(request.Page, request.PerPage, total)

	// list
	tags, err := s.Repository.List(ctx, pg.PerPage, pg.From, request.Keyword)
	if err != nil {
		return nil, err
	}
	var items []*tagpb.Tag
	for _, v  := range tags {
		item := &tagpb.Tag{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Status:      v.Status,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
		items = append(items, item)
	}

	return &tagpb.ListResponse{
		Total:       pg.Total,
		PerPage:     pg.PerPage,
		CurrentPage: pg.CurrentPage,
		LastPage:    pg.LastPage,
		From:        pg.From,
		To:          pg.To,
		Data:        items,
	}, nil
}
