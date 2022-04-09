package repository

import (
	"context"
	"database/sql"
	"time"
)

type UserTag struct {
	Id int64
	UserId int64
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateUserTagParam struct {
	Name string
}

type InsertUserTagParam struct {
	UserId int64
	Name string
}

const (
	sqlCreateUserTag = `
INSERT INTO user_tags(user_id, name, created_at, updated_at)
VALUES (?, ?, ?, ?)
`
	sqlUpdateUserTag = `
UPDATE user_tags SET name = ?, updated_at = ?
WHERE id = ? AND user_id = ?
`
	sqlDeleteUserTag = `
UPDATE user_tags SET deleted_at = ?
WHERE id = ? AND user_id = ?
`
	sqlGetUserTag = `
SELECT id, user_id, name, created_at, updated_at 
FROM user_tags
WHERE id=? AND user_id = ? AND deleted_at IS NULL
`
	sqlListUserTagCount = `
SELECT count(*)
FROM user_tags
WHERE user_id = ? AND deleted_at IS NULL
`
	sqlListUserTag = `
SELECT id, user_id, name, created_at, updated_at
FROM user_tags
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY created_at desc
LIMIT ? OFFSET ?
`
)

func (r *Repository) InsertUserTag(ctx context.Context, p *InsertUserTagParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlCreateUserTag)
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

func (r *Repository) UpdateUserTag(ctx context.Context, id, userId int64, p *UpdateUserTagParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlUpdateUserTag)
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

func (r *Repository) DeleteUserTag(ctx context.Context, id, userId int64) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlDeleteUserTag)
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

func (r *Repository) GetUserTag(ctx context.Context, id, userId int64) (*UserTag, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlGetUserTag)
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


func (r *Repository) ListUserTagCount(ctx context.Context, userId int64) (int64, error) {
	// count
	stmt, err := r.db.PrepareContext(ctx, sqlListUserTagCount)
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

func (r *Repository) ListUserTag(ctx context.Context, userId, limit, offset int64) ([]*UserTag, error) {
	// query
	stmt, err := r.db.PrepareContext(ctx, sqlListUserTag)
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
