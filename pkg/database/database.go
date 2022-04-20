package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"goblog.com/pkg/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	DB *sql.DB
	Gorm *gorm.DB
}

type Config struct {
	Driver string
	Host string
	Port string
	Username string
	Password string
	Database string
}

type Options struct {
	Env environment.Environment
	ConnMaxLifetime time.Duration
	MaxIdleConns int
	MaxOpenConns int
}

type Option func(*Options)

var defaultOption = Options{
	Env: environment.PRODUCTION,
	ConnMaxLifetime: time.Minute * 3,
	MaxIdleConns: 10,
	MaxOpenConns: 100,
}

func WithEnv(e environment.Environment) Option {
	return func(l *Options) {
		l.Env = e
	}
}

func WithConnMaxLifetime(t time.Duration) Option {
	return func(c *Options) {
		c.ConnMaxLifetime = t
	}
}

func WithMaxIdleConns(n int) Option {
	return func(c *Options) {
		c.MaxIdleConns = n
	}
}

func WithMaxOpenConns(n int) Option {
	return func(c *Options) {
		c.MaxOpenConns = n
	}
}

// NewDatabase
// dns refer https://github.com/go-sql-driver/mysql for details
func NewDatabase(c Config, opts ...Option) (*Database, error){
	if c.Driver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&multiStatements=true", c.Username, c.Password, c.Host, c.Port, c.Database)
		db, err := sql.Open(c.Driver, dsn)
		if err != nil {
			return nil, err
		}

		c := defaultOption
		for _, o := range opts {
			o(&c)
		}

		db.SetConnMaxLifetime(c.ConnMaxLifetime)
		db.SetMaxIdleConns(c.MaxIdleConns)
		db.SetMaxOpenConns(c.MaxOpenConns)

		// gorm
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: db,
		}), &gorm.Config{})

		return &Database{
			DB:   db,
			Gorm: gormDB,
		}, nil
	}
	return nil, errors.New("database: driver not support")

}

var ProviderSet = wire.NewSet(NewDatabase)