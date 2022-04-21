package application

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/pkg/database"
	"goblog.com/service/tag/internal/config"
)

type Application struct {
	Config *config.Config
	Database *database.Database
	Logger  *zap.Logger
}

func NewApplication(c *config.Config, d *database.Database, l *zap.Logger) (*Application, error) {
	return &Application{
		Config: c,
		Database: d,
		Logger: l,
	}, nil
}

var ProviderSet = wire.NewSet(NewApplication)