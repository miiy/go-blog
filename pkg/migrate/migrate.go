package migrate

import (
	pkgMigrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source"
)

type migrate struct {
	migrate *pkgMigrate.Migrate
}

// NewMigrate
// mysql dsn parameters multiStatements=true
func NewMigrate(sourceName string, sourceInstance source.Driver, databaseName string, databaseInstance database.Driver) (*migrate, error) {
	m, err := pkgMigrate.NewWithInstance(sourceName, sourceInstance, databaseName, databaseInstance)
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
