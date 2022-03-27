package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goblog.com/pkg/enviroment"
)

type Options struct {
	Env string
	Debug bool
}

type Callback func(r *gin.Engine)

func NewOptions() *Options {
	return &Options{
		Env:   "",
		Debug: false,
	}
}
func NewGin(o *Options) (*gin.Engine, error) {
	if o.Debug {
		fmt.Printf( "[App-debug] Router env is %s\n", o.Env)
	}
	if o.Env == enviroment.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else if o.Env == enviroment.TESTING {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	return engine, nil
}

var ProviderSet = wire.NewSet(NewOptions, NewGin)

//
//func (r *router) Register(callback Callback) {
//	callback(r)
//}
