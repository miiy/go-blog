// +build wireinject

package main

import (
	"github.com/google/wire"
	"goblog.com/pkg/database"
	"goblog.com/pkg/environment"
	"goblog.com/pkg/logger"
	"goblog.com/service/tag/internal/application"
	"goblog.com/service/tag/internal/config"
)

func InitApplication(conf string) (*application.Application, func(), error) {
	panic(wire.Build(
		config.ProviderSet,
		logger.ProviderSet, providerLoggerOption,
		providerDatabase, providerDatabaseOption, database.NewDatabase,
		application.ProviderSet,
	))
}


func providerLoggerOption(config *config.Config) []logger.Option {
	return []logger.Option{
		logger.WithEnv(environment.Environment(config.App.Env)),
	}
}

func providerDatabase(config *config.Config) database.Config {
	return database.Config{
		Driver:   config.Database.Driver,
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Database: config.Database.Database,
	}
}

func providerDatabaseOption(config *config.Config) []database.Option {
	return []database.Option{
		database.WithEnv(environment.Environment(config.App.Env)),
	}
}