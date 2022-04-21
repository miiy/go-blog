package repository

import (
	"context"
	"database/sql"
	"errors"
	"goblog.com/pkg/jwtauth"
	"strings"
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

type User struct {
	Id int64
	Username string
	Password string
	Status int64
}

type RegisterParam struct {
	Username string
	Password string
	PasswordConfirmation string
	Email string
}

const (
	sqlRegister = `INSERT INTO users(username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	sqlUserExist = `SELECT COUNT(*) FROM users WHERE username = ? OR email = ?`
	sqlFirstByUserName = `SELECT id, username, password, status FROM users WHERE username = ?`
	sqlFirstById = `SELECT id, username, password, status FROM users WHERE id = ?`

	FieldUsername = "username"
	FieldEmail = "email"
	FieldPhone = "phone"
	sqlFieldExist = `SELECT COUNT(*) FROM users where {field} = ?`
)

var (
	ErrFieldError = errors.New("field error")
)

func (r *Repository) SignUp(ctx context.Context, p RegisterParam) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlRegister)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()
	res, err := stmt.ExecContext(ctx, p.Username, p.Password, p.Email, now, now)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repository) UserExist(ctx context.Context, p RegisterParam) (bool, error) {
		stmt, err := r.db.PrepareContext(ctx, sqlUserExist)
		if err != nil {
			return false, err
		}
		defer stmt.Close()

		var count int64
		if err = stmt.QueryRowContext(ctx, p.Username, p.Email).Scan(&count); err != nil {
			return false, err
		}

		return count > 0, nil
}

func (r *Repository) FirstByUsername(ctx context.Context, username string) (*User, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlFirstByUserName)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user User

	err = stmt.QueryRowContext(ctx, username).Scan(&user.Id, &user.Username, &user.Password, &user.Status)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FirstById(ctx context.Context, id int64) (*User, error) {
	stmt, err := r.db.PrepareContext(ctx, sqlFirstById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user User

	err = stmt.QueryRowContext(ctx, id).Scan(&user.Id, &user.Username, &user.Password, &user.Status)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FieldExist(ctx context.Context, field string, value string) (bool, error) {
	if field != FieldUsername && field != FieldEmail && field != FieldPhone {
		return false, ErrFieldError
	}
	sqlString := strings.Replace(sqlFieldExist, "{field}", field, 1)

	stmt, err := r.db.PrepareContext(ctx, sqlString)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int64
	if err = stmt.QueryRowContext(ctx, value).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) RetrieveByUsername(ctx context.Context, username string) (*jwtauth.AuthUser, error) {
	user, err := r.FirstByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &jwtauth.AuthUser{
		Id:       user.Id,
		Username: user.Username,
	}, nil
}