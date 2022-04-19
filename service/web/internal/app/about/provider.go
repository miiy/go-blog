package about

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"goblog.com/service/web/internal/pkg/application"

	//"goblog.com/pkg/application"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module struct {
	router *gin.Engine
	database *gorm.DB
	redis redis.UniversalClient
	logger *zap.Logger
}

var module *Module

func NewModule(router *gin.Engine, database *gorm.DB, redis redis.UniversalClient, logger *zap.Logger) *Module {
	module = &Module{
		router: router,
		database: database,
		redis: redis,
		logger: logger,
	}
	return module
}

var ProviderSet = wire.NewSet(NewModule, wire.FieldsOf(new(*application.Application), "Database", "Redis", "Router", "Logger"))
