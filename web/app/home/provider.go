package home

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Module struct {
	router *gin.Engine
	logger *zap.Logger
}

var module *Module

func NewModule(router *gin.Engine, logger *zap.Logger) *Module {
	module = &Module{
		router: router,
		logger: logger,
	}
	return module
}

var ProviderSet = wire.NewSet(NewModule)
