package repository

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Feedback struct {
	Id int64
	UserId int64
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type InsertParam struct {
	UserId int64
	Content string
}

type UpdateParam struct {
	Content string
}

const (
	sqlCreate = `INSERT INTO feedbacks(user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?)`
	sqlSoftDelete = `UPDATE feedbacks SET deleted_at = ? WHERE id = ?`
	sqlListCount = `SELECT count(*) FROM feedbacks WHERE deleted_at IS NULL`
	sqlList = `
SELECT id, user_id, content, created_at, updated_at FROM feedbacks
WHERE deleted_at IS NULL
ORDER BY created_at desc
LIMIT ? OFFSET ?
`

)

func (r *Repository) Insert(ctx context.Context, p *InsertParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlCreate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()
	res, err := stmt.ExecContext(ctx, p.UserId, p.Content, now, now)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repository) Delete(ctx context.Context, id int64) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlSoftDelete)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx , time.Now(), id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (r *Repository) ListCount(ctx context.Context) (int64, error) {
	// count
	stmt, err := r.db.PrepareContext(ctx, sqlListCount)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var count int64
	if err = stmt.QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) List(ctx context.Context, userId, limit, offset int64) ([]*Feedback, error) {
	// query
	stmt, err := r.db.PrepareContext(ctx, sqlList)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var items []*Feedback
	rows, err := stmt.QueryContext(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i = new(Feedback)
		err = rows.Scan(
			&i.Id,
			&i.UserId,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		);
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

