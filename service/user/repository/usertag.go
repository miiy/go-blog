package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserTagRepository struct {
	db *sql.DB
}

func NewUserTagRepository(db *sql.DB) *UserTagRepository {
	return &UserTagRepository{
		db: db,
	}
}

type UserTag struct {
	Id int64
	UserId int64
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateParam struct {
	Name string
}

type InsertParam struct {
	UserId int64
	Name string
}

const (
	sqlCreate = `
INSERT INTO user_tags(user_id, name, created_at, updated_at)
VALUES (?, ?, ?, ?)
`
	sqlUpdate = `
UPDATE user_tags SET name = ?, updated_at = ?
WHERE id = ? AND user_id = ?
`
	sqlDelete = `
UPDATE user_tags SET deleted_at = ?
WHERE id = ? AND user_id = ?
`
	sqlGet = `
SELECT id, user_id, name, created_at, updated_at 
FROM user_tags
WHERE id=? AND user_id = ? AND deleted_at IS NULL
`
	sqlListCount = `
SELECT count(*)
FROM user_tags
WHERE user_id = ? AND deleted_at IS NULL
`
	sqlList = `
SELECT id, user_id, name, created_at, updated_at
FROM user_tags
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY created_at desc
LIMIT ? OFFSET ?
`
)

var ErrNotFound = errors.New("not found")

func (r *UserTagRepository) Insert(ctx context.Context, p *InsertParam) (int64, error) {
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

func (r *UserTagRepository) Update(ctx context.Context, id, userId int64, p *UpdateParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlUpdate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		p.Name,
		time.Now(),
		id,
		userId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *UserTagRepository) Delete(ctx context.Context, id, userId int64) (int64, error) {
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

func (r *UserTagRepository) Get(ctx context.Context, id, userId int64) (*UserTag, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlGet)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var i UserTag

	err = stmt.QueryRowContext(ctx, id, userId).Scan(
		&i.Id,
		&i.UserId,
		&i.Name,
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


func (r *UserTagRepository) ListCount(ctx context.Context, userId int64) (int64, error) {
	// count
	stmt, err := r.db.PrepareContext(ctx, sqlListCount)
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

func (r *UserTagRepository) List(ctx context.Context, userId, limit, offset int64) ([]*UserTag, error) {
	// query
	stmt, err := r.db.PrepareContext(ctx, sqlList)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var items []*UserTag
	rows, err := stmt.QueryContext(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i = new(UserTag)
		if err = rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Name,
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
