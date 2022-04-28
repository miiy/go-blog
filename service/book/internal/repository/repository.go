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
	CreateBook(ctx context.Context, book *Book) (*Book, error)
	UpdateBook(ctx context.Context, id int64, book *Book, selects interface{}) (*Book, error)
	DeleteBookById(ctx context.Context, id int64) error
	GetBookById(ctx context.Context, id int64) (*Book, error)
	FindCount(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
	Find(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*Book, error)
}

type Book struct {
	Id              int64
	UserId          int64
	CategoryId      int64
	Name            string
	Publisher       string
	Year            int
	Pages           int
	Price           float32
	Binding         string
	Series          string
	ISBN            string
	BookDescription string
	AboutTheAuthor  string
	TableOfContents string
	Content         string
	Status          int
	CreateTime      time.Time
	UpdateTime      time.Time
	DeleteTime      sql.NullTime
}

type BookMeta struct {
	Id              int64
	BookId          int64
	MetaTitle       string
	MetaDescription string
	CreateTime       time.Time `gorm:"autoCreateTime"`
	UpdateTime       time.Time `gorm:"autoUpdateTime"`
	DeleteTime       sql.NullTime
}

const (
	BookStatusDefault = 0
	BookStatusActive = 1
	BookStatusDisable = 2
)

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

func ScopeOfBookUser(userId int64) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId)
	}
}

func ScopeOfBookCategory(categoryId int64) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryId)
	}
}

func ScopeOfBookStatus(status []int64) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if len(status) == 1 {
			db.Where("status = ?", status)
		}
		return db.Where("status IN (?)", status)
	}
}

func ScopeBookActive() func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", BookStatusDisable)
	}
}

func ScopeBookDisable() func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", BookStatusDisable)
	}
}

func Paginate(page, pageSize int) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (r *BookRepositoryImpl) CreateBook(ctx context.Context, i *Book) (*Book, error) {
	err := r.db.WithContext(ctx).Create(&i).Error
	if err != nil {
		r.logger.Error(ErrCreateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return i, nil
}


func (r *BookRepositoryImpl) UpdateBook(ctx context.Context, id int64, i *Book, selects interface{}) (*Book, error) {
	err := r.db.WithContext(ctx).Model(&Book{}).Where("id = ?", id).Select(selects).Updates(i).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return i, nil
}

func (r *BookRepositoryImpl) DeleteBookById(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&Book{}, id).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return ErrUpdateError
	}
	return nil
}

func (r *BookRepositoryImpl) GetBookById(ctx context.Context, id int64) (*Book, error) {
	var i Book
	err := r.db.WithContext(ctx).Model(&Book{}).Scopes(ScopeBookActive()).First(&i, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	}
	return &i, nil
}

func (r *BookRepositoryImpl) FindCount(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	r.db.WithContext(ctx).Model(&Book{}).Scopes(scopes...).Count(&count)
	return count, nil
}

func (r *BookRepositoryImpl) Find(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*Book, error) {
	var items []*Book
	r.db.WithContext(ctx).Model(&Book{}).Scopes(scopes...).Find(&items)

	return items, nil
}
