package repository

import (
	"context"
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type BookRepositoryImpl struct {
	db *gorm.DB
	logger *zap.Logger
}

type BookRepository interface {
	Create(ctx context.Context, book *Book) (*Book, error)
	Update(ctx context.Context, id int64, book *Book) (*Book, error)
	Delete(ctx context.Context, id int64) error
	First(ctx context.Context, id int64, columns interface{}) (*Book, error)
	Find(ctx context.Context, limit, offset int64) ([]*Book, error)
}

type Book struct {
	Id              int64
	UserId          int64
	CategoryId      int64
	Name            string
	Publisher       string
	Year            int
	Pages           int
	Price           float64
	Binding         string
	series          string
	ISBN            string
	BookDescription string
	AboutTheAuthor  string
	TableOfContents string
	Content         string
	Status          int
	MetaTitle       string
	MetaDescription string
	CreatedTime     time.Time
	UpdatedTime     time.Time
	DeletedTime     sql.NullString
}

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrCreateError    = errors.New("create error")
	ErrUpdateError    = errors.New("update error")
)

func NewBookRepository(db *gorm.DB, logger *zap.Logger) BookRepository {
	return &BookRepositoryImpl{
		db: db,
		logger: logger,
	}
}

func (r *BookRepositoryImpl) Create(ctx context.Context, i *Book) (*Book, error) {
	err := r.db.WithContext(ctx).Create(&i).Error
	if err != nil {
		r.logger.Error(ErrCreateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return i, nil
}


func (r *BookRepositoryImpl) Update(ctx context.Context, id int64, i *Book) (*Book, error) {
	err := r.db.WithContext(ctx).Model(&Book{}).Where("id = ?", id).Updates(i).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return i, nil
}

func (r *BookRepositoryImpl) Delete(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&Book{}, id).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return ErrUpdateError
	}
	return nil
}

func (r *BookRepositoryImpl) First(ctx context.Context, id int64, columns interface{}) (*Book, error) {
	var i Book
	err := r.db.WithContext(ctx).Model(&Book{}).Select(columns).First(&i, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	}
	return &i, nil
}

func (r *BookRepositoryImpl) FindCount(ctx context.Context) (int64, error) {
	var count int64
	r.db.WithContext(ctx).Model(&Book{}).Count(&count)
	return count, nil
}

func (r *BookRepositoryImpl) Find(ctx context.Context, limit, offset int64) ([]*Book, error) {
	var items []*Book
	r.db.WithContext(ctx).Model(&Book{}).Find(&items).Limit(int(limit)).Offset(int(offset))

	return items, nil
}
