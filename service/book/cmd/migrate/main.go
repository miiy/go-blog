package main

import (
	"database/sql"
	"flag"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "goblog.com/pkg/config"
	"goblog.com/pkg/migrate"
	"goblog.com/service/book/migrations"
	"log"
	"strings"
)

const MigrationsTable = "book_migrations"

func main() {
	conf := flag.String("c", "./configs/default.yaml", "config file")
	cmd := flag.String("cmd", "up", "up, down")
	flag.Parse()

	app, cleanUp, err := InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()


	if err = run(app.Database.DB, app.Config.Database.Database, *cmd); err != nil {
		log.Fatalln(err)
	}
}

func run(db *sql.DB, dbName, cmd string) error {

	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return err
	}


	mDriver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{
		MigrationsTable: MigrationsTable,
	})
	if err != nil {
		return err
	}


	m, err := migrate.NewMigrate("s", d, dbName, mDriver)
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
		err = m.Down()
		if err != nil {
			return err
		}
		log.Println("success")
	}

	return nil
}
