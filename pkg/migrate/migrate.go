package migrate

import (
	pkg_migrate"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"
)

type migrate struct {
	migrate *pkg_migrate.Migrate
}

// NewMigrate
// set mysql dsn parameters multiStatements=true
func NewMigrate(sourceName string, sourceInstance source.Driver, databaseName string, databaseInstance database.Driver) (*migrate, error) {
	m, err := pkg_migrate.NewWithInstance(sourceName, sourceInstance, databaseName, databaseInstance)
	if err != nil {
		return nil, err
	}

	return &migrate{
		migrate: m,
	}, nil
}

func (m *migrate) Up() error {
	return m.migrate.Up()
}

func (m *migrate) Down() error {
	return m.migrate.Down()
}