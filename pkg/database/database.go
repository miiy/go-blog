package database

import (
	"database/sql"
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

type DSNString string

//type Options struct {
//	Driver string
//	Host string
//	Port string
//	Username string
//	Password string
//	Database string
//	Charset string
//}

type Config struct {
	Env environment.Environment
	ConnMaxLifetime time.Duration
	MaxIdleConns int
	MaxOpenConns int
}

type Option func(*Config)

var defaultConfig = Config{
	Env: environment.PRODUCTION,
	ConnMaxLifetime: time.Minute * 3,
	MaxIdleConns: 10,
	MaxOpenConns: 100,
}

func WithEnv(e environment.Environment) Option {
	return func(l *Config) {
		l.Env = e
	}
}

func WithConnMaxLifetime(t time.Duration) Option {
	return func(c *Config) {
		c.ConnMaxLifetime = t
	}
}

func WithMaxIdleConns(n int) Option {
	return func(c *Config) {
		c.MaxIdleConns = n
	}
}

func WithMaxOpenConns(n int) Option {
	return func(c *Config) {
		c.MaxOpenConns = n
	}
}

// NewDatabase
// dns refer https://github.com/go-sql-driver/mysql for details
func NewDatabase(dsn DSNString, opts ...Option) (*Database, error){
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&multiStatements=true", o.Username, o.Password, o.Host, o.Port, o.Database, o.Charset)
	db, err := sql.Open("mysql", string(dsn))
	if err != nil {
		return nil, err
	}

	c := defaultConfig
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

func ProviderDSNString(s string) DSNString {
	return DSNString(s)
}


var ProviderSet = wire.NewSet(NewDatabase)