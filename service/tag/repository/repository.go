package repository

import (
	"context"
	"database/sql"
	"errors"
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

const (
	statusDefault = 0
	statusActive = 1
	statusDisable = 2
)

type Tag struct {
	Id int64
	Name string
	Description string
	Status int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type InsertParam struct {
	Name string
	Description string
	Status int64
}

type UpdateParam struct {
	Name string
	Description string
	Status int64
}

const (
	sqlCreate = `INSERT INTO tags(name, description, created_at, updated_at) VALUES (?, ?, ?, ?)`
	sqlUpdate = `UPDATE tags SET name = ?, description = ?, updated_at = ? WHERE id = ?`
	sqlSoftDelete = `UPDATE tags SET deleted_at = ? WHERE id = ?`
	sqlGet = `SELECT id, name, description, created_at, updated_at FROM tags WHERE id=? AND deleted_at IS NULL`
	sqlListCount = `SELECT count(*) FROM tags WHERE deleted_at IS NULL`
	sqlListWithSearchCount =`SELECT count(*) FROM tags WHERE keyword LIKE '%?%' AND deleted_at IS NULL`
	sqlList = `
SELECT id, name, description, status, created_at, updated_at FROM tags
WHERE deleted_at IS NULL
ORDER BY created_at desc
LIMIT ? OFFSET ?
`
	sqlListWithSearch = `
SELECT id, name, description, created_at, updated_at FROM tags
WHERE keyword LIKE '%?%' AND deleted_at IS NULL
ORDER BY created_at desc
LIMIT ? OFFSET ?
`
)

var ErrorNotFound = errors.New("not found")

func (r *Repository) Insert(ctx context.Context, p *InsertParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlCreate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()
	res, err := stmt.ExecContext(ctx, p.Name, p.Description, now, now)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repository) Update(ctx context.Context, id int64, p *UpdateParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlUpdate)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.Name, p.Description, time.Now(), id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
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

func (r *Repository) Get(ctx context.Context, id int64) (*Tag, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlGet)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var i Tag

	err = stmt.QueryRowContext(ctx, id).Scan(&i.Id, &i.Name, &i.Description, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &i, nil
}

func (r *Repository) ListCount(ctx context.Context, keyword string) (int64, error) {
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
	if err = stmt.QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int64, keyword string) ([]*Tag, error) {
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

	var items []*Tag
	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var i = new(Tag)
		err = rows.Scan(
			&i.Id,
			&i.Name,
			&i.Description,
			&i.Status,
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

