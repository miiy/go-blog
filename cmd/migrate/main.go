package main

import (
	"flag"
	m_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	go_blog "github.com/miiy/go-blog"
	"github.com/miiy/go-blog/pkg/config"
	"github.com/miiy/go-blog/pkg/database"
	"github.com/miiy/go-blog/pkg/migrate"
	"log"
)

func main() {
	cFile := flag.String("c", "./config/default.yaml", "config path")

	flag.Parse()
	if err := execute(*cFile); err != nil {
		log.Fatalln(err)
	}

}

func execute(conf string) error {
	// config
	c, err := config.NewConfig(conf)
	if err != nil {
		log.Fatalln(err)
	}

	// db
	db, err := database.NewDatabase(&c.Database)
	if err != nil {
		log.Fatalln(err)
	}

	d, err := iofs.New(go_blog.MigrationFS, "migrations")
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := db.Gorm.DB()
	if err != nil {
		log.Fatalln(err)
	}

	mDriver, err := m_mysql.WithInstance(sqlDB, &m_mysql.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewMigrate("s", d, c.Database.Database, mDriver)
	if err != nil {
		log.Fatalf("failed to migrate: %v\n", err)
	}


	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}