package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goblog.com/pkg/environment"
)

type Config struct {
	Env   environment.Environment
}

type Option func(*Config)

var defaultConfig = Config{
	Env: environment.DEVELOPMENT,
}

func WithEnv(e environment.Environment) Option {
	return func(o *Config) {
		o.Env = e
	}
}

func NewGin(opts ...Option) (*gin.Engine, error) {
	conf := defaultConfig
	for _, o := range opts {
		o(&conf)
	}
	if conf.Env == environment.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else if conf.Env == environment.TESTING {
		gin.SetMode(gin.TestMode)
	} else {
		fmt.Printf( "[App-debug] Router env is %s\n", conf.Env)
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	return engine, nil
}

var ProviderSet = wire.NewSet(NewGin)
