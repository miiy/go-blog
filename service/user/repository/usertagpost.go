package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserTagPostRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserTagPostRepository {
	return &UserTagPostRepository{
		db: db,
	}
}

type UserTagPost struct {
	Id int64
	UserId int64
	TagId int64
	PostId int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InsertParam struct {
	UserId int64
	Name string
}

const (
	sqlCreate = `
INSERT INTO user_tags_posts(user_id, tag_id, post_id, created_at, updated_at)
VALUES (?, ?, ?, ?)
`
	sqlDelete = `
UPDATE user_tags_posts SET deleted_at = ?
WHERE id = ? AND user_id = ?
`
	sqlPostList = `
SELECT p.user_id, p.title, p.content, p.status, p.published_time, p.updated_time, p.sort
FROM user_tags_posts as tp JOIN user_posts as p ON tp.post_id = p.id
WHERE user_id = ? AND tag_id = ? AND deleted_at IS NULL
ORDER BY published_time DESC
LIMIT ? OFFSET ?
`
	sqlPostListCount = `
SELECT count(*)
FROM user_tags_posts as tp JOIN user_posts as p ON tp.post_id = p.id
WHERE user_id = ? AND tag_id = ? AND deleted_at IS NULL
`
)

var ErrNotFound = errors.New("not found")

func (r *UserTagPostRepository) Insert(ctx context.Context, p *InsertParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlCreate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()
	res, err := stmt.ExecContext(
		ctx,
		p.UserId,
		p.Name,
		now,
		now,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *UserTagPostRepository) Delete(ctx context.Context, id, userId int64) (int64, error) {
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

func (r *UserTagPostRepository) PostListCount(ctx context.Context, userId, tagId int64) (int64, error) {
	// count
	stmt, err := r.db.PrepareContext(ctx, sqlPostListCount)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int64
	if err = stmt.QueryRowContext(ctx, userId, tagId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserTagPostRepository) PostList(ctx context.Context, userId, tagId, limit, offset int64) ([]*UserTagPost, error) {
	// query
	stmt, err := r.db.PrepareContext(ctx, sqlPostList)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var items []*UserTagPost
	rows, err := stmt.QueryContext(ctx, userId, tagId, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i = new(UserTagPost)
		if err = rows.Scan(
			&i.Id,
			&i.UserId,
			&i.TagId,
			&i.PostId,
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
