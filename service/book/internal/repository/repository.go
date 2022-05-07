package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"goblog.com/pkg/database/gorm/model"
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
	model.Model
	UserId          int64        `gorm:"column:user_id"`
	CategoryId      int64        `gorm:"column:category_id"`       //分类
	Name            string       `gorm:"column:name"`              //书名
	Publisher       string       `gorm:"column:publisher"`         //出版社
	Year            int          `gorm:"column:year"`              //出版年
	Pages           int          `gorm:"column:pages"`             //页数
	Price           float32      `gorm:"column:price"`             //定价
	Binding         string       `gorm:"column:binding"`           //装帧
	Series          string       `gorm:"column:series"`            //丛书
	ISBN            string       `gorm:"column:isbn"`              //ISBN
	BookDescription string       `gorm:"column:book_description"`  //图书简介
	AboutTheAuthor  string       `gorm:"column:about_the_author"`  //作者简介
	TableOfContents string       `gorm:"column:table_of_contents"` //目录
	Content         string       `gorm:"column:content"`           //内容
	Status          int          `gorm:"column:status"`
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
	BookFieldNames string
	BookFieldNamesExpectAutoSet string
	BookFieldNamesWithPlaceHolder string
)

var(
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrCreateError    = errors.New("create error")
	ErrUpdateError    = errors.New("update error")
)

func NewBookRepository(db *gorm.DB, logger *zap.Logger) BookRepository {
	bookFields, err := model.FieldDBNames(&Book{}, nil)
	if err != nil {
		logger.Error(err.Error())
	}
	bookFieldsExceptAutoSet, err := model.FieldDBNames(&Book{}, model.FieldNameExpectAutoSet)
	if err != nil {
		logger.Error(err.Error())
	}

	BookFieldNames = model.FieldNameFormat(bookFields, model.FieldNameFormatWithQuote)
	fmt.Println(BookFieldNames)
	BookFieldNamesExpectAutoSet = model.FieldNameFormat(bookFieldsExceptAutoSet, model.FieldNameFormatWithQuote)
	fmt.Println(BookFieldNamesExpectAutoSet)
	BookFieldNamesWithPlaceHolder = model.FieldNameFormat(bookFieldsExceptAutoSet, model.FieldNameFormatWithPlaceHolder)
	fmt.Println(BookFieldNamesWithPlaceHolder)

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

func WithSelect(s string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Select(s)
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
