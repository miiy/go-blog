package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	DB *sql.DB
	Gorm *gorm.DB
}

type Options struct {
	Driver string
	Host string
	Port string
	Username string
	Password string
	Database string
	Charset string
}

func NewDatabase(o *Options) (*Database, error){
	if o.Driver == "mysql" {
		// refer https://github.com/go-sql-driver/mysql for details
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&multiStatements=true", o.Username, o.Password, o.Host, o.Port, o.Database, o.Charset)
		db, err := sql.Open(o.Driver, dsn)
		if err != nil {
			return nil, err
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)

		// gorm
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: db,
		}), &gorm.Config{})

		return &Database{
			DB:   db,
			Gorm: gormDB,
		}, nil
	}

	return nil, errors.New("driver not support")
}
