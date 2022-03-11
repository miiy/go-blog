package logger

import (
	"go.uber.org/zap"
)

type Options struct {
	Env string
}

func NewLogger(o *Options) (*zap.Logger, func(), error) {
	var logger *zap.Logger
	var err error

	if o.Env == "local" || o.Env == "testing" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, nil, err
	}

	cleanUp := func() {
		defer logger.Sync()
	}
	return logger, cleanUp, nil
}
