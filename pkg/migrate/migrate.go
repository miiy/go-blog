package migrate

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(db *sql.DB, sourceURL string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "mysql", driver)
	if err != nil {
		return err
	}
	if err = m.Steps(1); err != nil {
		return err
	}
	return nil
}