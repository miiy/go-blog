package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	Content string
	Status int64
	PublishedTime time.Time
	UpdatedTime time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type APIArticle struct {
	Id int64
	UserId int64
	Title string
	Content string
	Status int
}

type InsertParam struct {
	UserId int64
	Title string
	Content string
	Status int64
	PublishedTime time.Time
	UpdatedTime time.Time
	Sort int64
}

type UpdateParam struct {
	Title string
	Content string
	Status int64
	PublishedTime time.Time
	UpdatedTime time.Time
	Sort int64
}

var ErrNotFound = errors.New("not found")

func (r *Repository) Create(ctx context.Context, p *InsertParam) (int64, error) {
	article := &Article{
		Id:            0,
		UserId:        p.UserId,
		Title:         p.Title,
		Content:       p.Content,
		Status:        p.Status,
		PublishedTime: p.PublishedTime,
		UpdatedTime:   p.UpdatedTime,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := r.db.Create(&article).Error
	return article.Id, err
}

func (r *Repository) Update(ctx context.Context, id, userId int64, p *UpdateParam) (int64, error) {
	article := Article{
		Title:         p.Title,
		Content:       p.Content,
		Status:        p.Status,
		PublishedTime: p.PublishedTime,
		UpdatedTime:   p.UpdatedTime,
	}
	result := r.db.Model(&Article{}).Where("id = ?", id).Updates(article)
	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(ctx context.Context, id, userId int64) (int64, error) {
	result := r.db.Delete(&Article{}, id)
	return result.RowsAffected, result.Error
}

func (r *Repository) First(ctx context.Context, id int64) (*APIArticle, error) {
	var apiArticle APIArticle
	r.db.Model(&Article{}).First(&apiArticle, id)
	return &apiArticle, nil
}


func (r *Repository) FindCount(ctx context.Context, userId int64, keyword string) (int64, error) {
	sqlStr := sqlListCount
	if keyword != "" {
		sqlStr = sqlListWithSearchCount
	}

	// count
	stmt, err := r.db.PrepareContext(ctx, sqlStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int64
	if err = stmt.QueryRowContext(ctx, userId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.Query("page_size"))
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

func (r *Repository) Find(ctx context.Context, userId, limit, offset int64, keyword string) ([]*UserPost, error) {
	var articles []APIArticle
	r.db.Model(&Article{}).Scopes(Paginate(r)).Find(&articles)

	return items, nil
}
