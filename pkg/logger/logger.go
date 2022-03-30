package logger

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"goblog.com/pkg/environment"
)

type Config struct {
	Env environment.Environment
}

type Option func(*Config)

var defaultConfig = Config{
	Env: environment.PRODUCTION,
}

func WithEnv(e environment.Environment) Option {
	return func(l *Config) {
		l.Env = e
	}
}

func NewLogger(opts ...Option) (*zap.Logger, func(), error) {
	c := defaultConfig
	for _, o := range opts {
		o(&c)
	}

	var logger *zap.Logger
	var err error

	if c.Env == environment.DEVELOPMENT || c.Env == environment.TESTING {
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


var ProviderSet = wire.NewSet(NewLogger)
