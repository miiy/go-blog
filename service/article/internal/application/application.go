package application

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/pkg/database"
	"goblog.com/service/article/internal/config"
)

type Application struct {
	Config *config.Config
	Database *database.Database
	Logger  *zap.Logger
}

func NewApplication(conf *config.Config, db *database.Database, logger *zap.Logger) *Application {
	return &Application{
		Config:   conf,
		Database: db,
		Logger:   logger,
	}
}

var ProviderSet = wire.NewSet(NewApplication)
