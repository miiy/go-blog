package main

import (
	"database/sql"
	"flag"
	"github.com/go-sql-driver/mysql"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "goblog.com/pkg/config"
	"goblog.com/pkg/migrate"
	mainpkg "goblog.com/service/article/cmd/article"
	"goblog.com/service/article/migrations"
	"log"
	"strings"
)

func main() {
	conf := flag.String("c", "./config/default.yaml", "config file")
	cmd := flag.String("cmd", "up", "up, down")
	flag.Parse()

	app, cleanUp, err := mainpkg.InitApplication(*conf)
	if err != nil {
		panic(err)
	}
	defer cleanUp()


	mysqlCfg, err := mysql.ParseDSN(app.Config.Mysql.DSN)
	if err != nil {
		panic(err)
	}

	if err := run(app.Database.DB, mysqlCfg.DBName, *cmd); err != nil {
		log.Fatalln(err)
	}
}

func run(db *sql.DB, dbName, cmd string) error {


	d, err := iofs.New(migrations.FS, "migrations")
	if err != nil {
		return err
	}


	mDriver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{})
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
		err = m.Up()
		if err != nil {
			return err
		}
		log.Println("success")
	}

	return nil
}
