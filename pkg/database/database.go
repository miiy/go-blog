package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Database struct {
	DB *sql.DB
}

type Options struct {
	Driver string
	Host string
	Port string
	Username string
	Password string
	Database string
}

func NewDatabase(o *Options) (*Database, error) {
	if o.Driver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", o.Username, o.Password, o.Host, o.Port, o.Database)
		db, err := sql.Open(o.Driver, dsn)
		if err != nil {
			return nil, err
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)

		return &Database{
			DB: db,
		}, nil
	}
	return nil, errors.New("driver not support")
}
