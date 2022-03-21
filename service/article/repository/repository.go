package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Article struct {
	Id int64
	UserId int64
	Title string
	MetaTitle string
	MetaDescription string
	PublishedTime time.Time
	UpdatedTime time.Time
	FromTest string
	FromUrl string
	Summary string
	Content string
	Status int
	CreatedAt time.Time
	UpdatedAt time.Time
}

var ErrRecordNotFound = gorm.ErrRecordNotFound

func (r *Repository) Create(ctx context.Context, a *Article) (*Article, error) {
	err := r.db.Create(&a).Error
	return a, err
}

func (r *Repository) Update(ctx context.Context, id uint64, a *Article) (*Article, error) {
	result := r.db.Model(&Article{}).Where("id = ?", id).Updates(a)
	return a, result.Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) (int64, error) {
	result := r.db.Delete(&Article{}, id)
	return result.RowsAffected, result.Error
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

func (r *Repository) Find(ctx context.Context, limit, offset int) ([]*Article, error) {
	var items []*Article
	r.db.Model(&Article{}).Find(&items).Limit(limit).Offset(offset)

	return items, nil
}
