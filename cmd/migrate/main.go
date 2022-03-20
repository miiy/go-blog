package main

import (
	"flag"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	//goBlog "github.com/miiy/go-blog"
	"goblog.com/pkg/config"
	"goblog.com/pkg/database"
	"goblog.com/pkg/migrate"
	"log"
	"strings"
)

func main() {
	cFile := flag.String("c", "./config/default.yaml", "config path")
	cmd := flag.String("cmd", "up", "up, down")
	flag.Parse()
	if err := run(*cFile, *cmd); err != nil {
		log.Fatalln(err)
	}
}

func run(conf, cmd string) error {
	// config
	c, err := config.NewConfig(conf)
	if err != nil {
		return err
	}

	// db
	db, err := database.NewDatabase(&c.Database)
	if err != nil {
		return err
	}

	d, err := iofs.New(goBlog.MigrationFS, "migrations")
	if err != nil {
		return err
	}

	sqlDB, err := db.Gorm.DB()
	if err != nil {
		return err
	}

	mDriver, err := migrateMysql.WithInstance(sqlDB, &migrateMysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewMigrate("s", d, c.Database.Database, mDriver)
	if err != nil {
		return err
	}

	if "up" == strings.ToLower(cmd) {
		err = m.Up()
		if err != nil {
			return err
		}
		log.Println("success")
	}

	if "down" == strings.ToLower(cmd) {
		err = m.Up()
		if err != nil {
			return err
		}
		log.Println("success")
	}

	return nil
}
