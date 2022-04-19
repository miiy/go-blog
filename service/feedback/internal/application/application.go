package application

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/pkg/database"
)

type Application struct {
	Database *database.Database
	Logger  *zap.Logger
}

func NewApplication(database *database.Database, logger *zap.Logger) (*Application, error) {
	return &Application{
		Database: database,
		Logger: logger,
	}, nil
}

var ProviderSet = wire.NewSet(NewApplication)