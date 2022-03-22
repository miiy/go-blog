package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRepository(db *gorm.DB, logger *zap.Logger) *Repository {
	return &Repository{
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

func (r *Repository) Create(ctx context.Context, a *Article) (*Article, error) {
	err := r.db.Create(&a).Error
	if err != nil {
		r.logger.Error(ErrCreateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return a, nil
}

func (r *Repository) Update(ctx context.Context, id uint64, a *Article) (*Article, error) {
	err := r.db.Model(&Article{}).Where("id = ?", id).Updates(a).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return nil, ErrUpdateError
	}
	return a, nil
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	err := r.db.Delete(&Article{}, id).Error
	if err != nil {
		r.logger.Error(ErrUpdateError.Error(), zap.Error(err))
		return ErrUpdateError
	}
	return nil
}

func (r *Repository) First(ctx context.Context, id uint64) (*Article, error) {
	var article Article
	err := r.db.Model(&Article{}).First(&article, id).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	}
	return &article, nil
}

func (r *Repository) FindCount(ctx context.Context) (int64, error) {
	var count int64
	r.db.Model(&Article{}).Count(&count)
	return count, nil
}

func (r *Repository) Find(ctx context.Context, limit, offset int32) ([]*Article, error) {
	var items []*Article
	r.db.Model(&Article{}).Find(&items).Limit(int(limit)).Offset(int(offset))

	return items, nil
}
