package server

import (
	"context"
	"go.uber.org/zap"
	bookpb "goblog.com/api/book/v1"
	"goblog.com/pkg/database/gorm/paginate"
	"goblog.com/service/book/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type BookServer struct {
	Repository repository.BookRepository
	Logger *zap.Logger
	bookpb.UnimplementedBookServiceServer
}

func NewBookServer(db *gorm.DB, logger *zap.Logger) bookpb.BookServiceServer {
	r := repository.NewBookRepository(db, logger)
	return &BookServer{
		Repository: r,
		Logger: logger,
	}
}

func (s *BookServer) CreateBook(ctx context.Context, request *bookpb.CreateBookRequest) (*bookpb.Book, error) {
	reqBook := request.GetBook()

	book := &repository.Book{
		UserId:          reqBook.UserId,
		CategoryId:      reqBook.CategoryId,
		Name:            reqBook.Name,
		Publisher:       reqBook.Publisher,
		Year:            int(reqBook.Year),
		Pages:           int(reqBook.Pages),
		Price:           reqBook.Price,
		Binding:         reqBook.Binding,
		ISBN:            reqBook.Isbn,
		BookDescription: reqBook.BookDescription,
		AboutTheAuthor:  reqBook.AboutTheAuthor,
		TableOfContents: reqBook.TableOfContents,
		Content:         reqBook.Content,
		Status:          int(reqBook.Status),
	}

	_, err := s.Repository.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}
	bookPb := bookToProto(book)
	return bookPb, nil
}

func (s *BookServer) BatchCreateBooks(ctx context.Context, request *bookpb.BatchCreateBooksRequest) (*bookpb.BatchCreateBooksResponse, error) {
	panic("implement me")
}

func (s *BookServer) GetBook(ctx context.Context, request *bookpb.GetBookRequest) (*bookpb.Book, error) {
	book, err := s.Repository.GetBookById(ctx, request.Id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			st := status.New(codes.NotFound, err.Error())
			return nil, st.Err()
		}
		return nil, err
	}

	return bookToProto(book), nil
}

func (s *BookServer) UpdateBook(ctx context.Context, request *bookpb.UpdateBookRequest) (*bookpb.Book, error) {
	book := request.GetBook()
	if ! request.UpdateMask.IsValid(book) {
		st := status.New(codes.InvalidArgument, "invalid field mask")
		return nil, st.Err()
	}

	selects := []string{"*"}
	if request.UpdateMask.GetPaths() != nil {
		selects = request.UpdateMask.GetPaths()
	}

	b := &repository.Book{
		UserId:          book.UserId,
		CategoryId:      book.CategoryId,
		Name:            book.Name,
		Publisher:       book.Publisher,
		Year:            int(book.Year),
		Pages:           int(book.Pages),
		Price:           book.Price,
		Binding:         book.Binding,
		Series:          book.Series,
		ISBN:            book.Isbn,
		BookDescription: book.BookDescription,
		AboutTheAuthor:  book.AboutTheAuthor,
		TableOfContents: book.TableOfContents,
		Content:         book.Content,
		Status:          int(book.Status),
	}
	ra, err := s.Repository.UpdateBook(ctx, request.Id, b, selects)
	if err != nil {
		return nil, err
	}

	return bookToProto(ra), nil
}

func (s *BookServer) DeleteBook(ctx context.Context, request *bookpb.DeleteBookRequest) (*emptypb.Empty, error) {
	err := s.Repository.DeleteBookById(ctx, request.Id)
	return nil, err
}

func (s *BookServer) ListBooks(ctx context.Context, request *bookpb.ListBooksRequest) (*bookpb.ListBooksResponse, error) {
	// count
	total, err := s.Repository.FindCount(ctx,
		repository.ScopeBookActive(),
		repository.ScopeOfBookCategory(request.CategoryId),
	)
	if err != nil {
		return nil, err
	}
	scopePaginate, totalPages := paginate.Paginate(int(request.Page), int(request.PageSize), int(total))

	// list
	books, err := s.Repository.Find(ctx,
		repository.ScopeBookActive(),
		repository.ScopeOfBookCategory(request.CategoryId),
		scopePaginate,
	)
	if err != nil {
		return nil, err
	}

	var items []*bookpb.Book
	for _, v  := range books {
		items = append(items, bookToProto(v))
	}

	return &bookpb.ListBooksResponse{
		Total:       total,
		TotalPages:  int64(totalPages),
		PageSize:    request.PageSize,
		CurrentPage: request.Page,
		Books:       items,
	}, nil
}

func bookToProto(book *repository.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:              book.Id,
		UserId:          book.UserId,
		CategoryId:      book.CategoryId,
		Name:            book.Name,
		Publisher:       book.Publisher,
		Year:            int64(book.Year),
		Pages:           int64(book.Pages),
		Price:           book.Price,
		Binding:         book.Binding,
		Series:          book.Series,
		Isbn:            book.ISBN,
		BookDescription: book.BookDescription,
		AboutTheAuthor:  book.AboutTheAuthor,
		TableOfContents: book.TableOfContents,
		Content:         book.Content,
		Status:          bookpb.Book_BookStatus(book.Status),
		CreateTime:      timestamppb.New(book.CreateTime),
		UpdateTime:      timestamppb.New(book.UpdateTime),
		DeleteTime:      timestamppb.New(book.DeleteTime.Time),
	}
}