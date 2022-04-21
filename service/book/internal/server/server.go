package server

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	bookpb "goblog.com/api/book/v1"
	"goblog.com/service/book/internal/repository"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
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
		Id:              reqBook.Id,
		UserId:          reqBook.UserId,
		CategoryId:      reqBook.CategoryId,
		Name:            reqBook.Name,
		Publisher:       reqBook.Publisher,
		Year:            int(reqBook.Year),
		Pages:           int(reqBook.Pages),
		Price:           float64(reqBook.Price),
		Binding:         reqBook.Binding,
		ISBN:            reqBook.Isbn,
		BookDescription: reqBook.BookDescription,
		AboutTheAuthor:  reqBook.AboutTheAuthor,
		TableOfContents: reqBook.TableOfContents,
		Content:         reqBook.Content,
		Status:          int(reqBook.Status),
		MetaTitle:       reqBook.MetaTitle,
		MetaDescription: reqBook.MetaDescription,
		CreatedTime:     time.Now(),
		UpdatedTime:     time.Now(),
		DeletedTime:     sql.NullTime{},
	}

	_, err := s.Repository.Create(ctx, book)
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
	panic("implement me")
}

func (s *BookServer) UpdateBook(ctx context.Context, request *bookpb.UpdateBookRequest) (*bookpb.Book, error) {
	panic("implement me")
}

func (s *BookServer) DeleteBook(ctx context.Context, request *bookpb.DeleteBookRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (s *BookServer) ListBooks(ctx context.Context, request *bookpb.ListBooksRequest) (*bookpb.ListBooksResponse, error) {
	panic("implement me")
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
		Price:           float32(book.Price),
		Binding:         book.Binding,
		Series:          book.Series,
		Isbn:            book.ISBN,
		BookDescription: book.BookDescription,
		AboutTheAuthor:  book.AboutTheAuthor,
		TableOfContents: book.TableOfContents,
		Content:         book.Content,
		Status:          bookpb.Book_BookStatus(book.Status),
		MetaTitle:       book.MetaTitle,
		MetaDescription: book.MetaDescription,
		CreateTime:      timestamppb.New(book.CreatedTime),
		UpdateTime:      timestamppb.New(book.UpdatedTime),
		DeletedTime:     timestamppb.New(book.DeletedTime.Time),
	}
}