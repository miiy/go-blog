package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type ArticleRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewArticleRepository(db *gorm.DB, logger *zap.Logger) *ArticleRepository {
	return &ArticleRepository{
		db:     db,
		logger: logger,
	}
}

type Article struct {
	Id              uint64
	UserId          uint64
	CategoryId      uint64
	Title           string
	MetaTitle       string
	MetaDescription string
	PublishedTime   time.Time
	UpdatedTime     time.Time
	FromText        string
	FromUrl         string
	Summary         string
	Content         string
	Status          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrCreateError    = errors.New("create error")
	ErrUpdateError    = errors.New("update error")
)

func (r *ArticleRepository) Create(ctx context.Context, a *Article) (*Article, error) {
	err := r.db.WithContext(ctx).Create(&a).Error
	if err != nil {
		r.logger.Error(ErrCreateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return a, nil
}

func (r *ArticleRepository) Update(ctx context.Context, id uint64, a *Article) (*Article, error) {
	err := r.db.WithContext(ctx).Model(&Article{}).Where("id = ?", id).Updates(a).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return a, nil
}

func (r *ArticleRepository) Delete(ctx context.Context, id uint64) error {
	err := r.db.WithContext(ctx).Delete(&Article{}, id).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return ErrUpdateError
	}
	return nil
}

func (r *ArticleRepository) First(ctx context.Context, id uint64, columns interface{}) (*Article, error) {
	var article Article
	err := r.db.WithContext(ctx).Model(&Article{}).Select(columns).First(&article, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	}
	return &article, nil
}

func (r *ArticleRepository) FindCount(ctx context.Context) (int64, error) {
	var count int64
	r.db.WithContext(ctx).Model(&Article{}).Count(&count)
	return count, nil
}

func (r *ArticleRepository) Find(ctx context.Context, limit, offset int64) ([]*Article, error) {
	var items []*Article
	r.db.WithContext(ctx).Model(&Article{}).Find(&items).Limit(int(limit)).Offset(int(offset))

	return items, nil
}
