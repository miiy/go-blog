package repository

import (
	"context"
	"database/sql"
	"errors"
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
	Content string
	Status int64
	PublishedTime time.Time
	UpdatedTime time.Time
	Sort int64
	CreatedAt time.Time
	UpdatedAt time.Time
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

const (
	sqlCreate = `
INSERT INTO user_posts(user_id, title, content, status, published_time, updated_time, sort, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`
	sqlUpdate = `
UPDATE user_posts
SET title = ?, content = ?, status = ?, published_time=?, updated_time=?, sort = ?, updated_at = ?
WHERE id = ? AND user_id = ?
`
	sqlUpdateSort = `
UPDATE user_posts
SET sort = ?, updated_at = ?
WHERE id = ? AND user_id = ?
`
	sqlDelete = `
UPDATE user_posts SET deleted_at = ?
WHERE id = ? AND user_id = ?
`
	sqlGet = `
SELECT id, user_id, title, content, status, published_time, updated_time, sort, created_at, updated_at 
FROM user_posts
WHERE id=? AND user_id = ? AND deleted_at IS NULL
`
	sqlListCount = `
SELECT count(*)
FROM user_posts
WHERE user_id = ? AND deleted_at IS NULL
`
	sqlListWithSearchCount = `
SELECT count(*)
FROM user_posts
WHERE user_id = ? AND title like '%?%' AND deleted_at IS NULL
`
	sqlList = `
SELECT id, user_id, title, content, status, published_time, updated_time, sort, created_at, updated_at
FROM user_posts
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY sort, published_time desc
LIMIT ? OFFSET ?
`
	sqlListWithSearch = `
SELECT id, user_id, title, content, status, published_time, updated_time, sort, created_at, updated_at
FROM user_posts
WHERE user_id = ? AND title like '%?%' AND deleted_at IS NULL
ORDER BY sort, published_time desc
LIMIT ? OFFSET ?
`
)

var ErrNotFound = errors.New("not found")

func (r *Repository) Insert(ctx context.Context, p *InsertParam) (int64, error) {
	now := time.Now()
	err := r.db.Create(Article{
		Id:            0,
		UserId:        p.UserId,
		Title:         p.Title,
		Content:       p.Content,
		Status:        p.Status,
		PublishedTime: p.PublishedTime,
		UpdatedTime:   p.UpdatedTime,
		Sort:          p.Sort,
		CreatedAt:     now,
		UpdatedAt:     now,
	}).Error
	if err != nil {
		return 0, err
	}

	return r.db.LastInsertId(), nil
}

func (r *Repository) Update(ctx context.Context, id, userId int64, p *UpdateParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlUpdate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.Status,
		p.PublishedTime,
		p.UpdatedTime,
		p.Sort,
		time.Now(),
		id,
		userId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Repository) Delete(ctx context.Context, id, userId int64) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlDelete)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx , time.Now(), id, userId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Repository) UpdateSort(ctx context.Context, id, userId, sort int64) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlUpdateSort)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx , sort, time.Now(), id, userId)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Repository) Get(ctx context.Context, id, userId int64) (*UserPost, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlGet)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var i UserPost

	err = stmt.QueryRowContext(ctx, id, userId).Scan(
		&i.Id,
		&i.UserId,
		&i.Title,
		&i.Content,
		&i.Status,
		&i.PublishedTime,
		&i.UpdatedTime,
		&i.Sort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &i, nil
}


func (r *Repository) ListCount(ctx context.Context, userId int64, keyword string) (int64, error) {
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

func (r *Repository) List(ctx context.Context, userId, limit, offset int64, keyword string) ([]*UserPost, error) {
	sqlStr := sqlList
	if keyword != "" {
		sqlStr = sqlListWithSearch
	}

	// query
	stmt, err := r.db.PrepareContext(ctx, sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var items []*UserPost
	rows, err := stmt.QueryContext(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i = new(UserPost)
		if err = rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Title,
			&i.Content,
			&i.Status,
			&i.PublishedTime,
			&i.UpdatedTime,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
