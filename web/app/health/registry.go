package health

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/miiy/go-web/pkg/application"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module struct {
	router *gin.Engine
	db *gorm.DB
	redis redis.UniversalClient
	logger *zap.Logger
}

var (
	module *Module
)

func NewModule(router *gin.Engine, database *gorm.DB, redis redis.UniversalClient, logger *zap.Logger) *Module {
	module = &Module{
		router: router,
		db: database,
		redis: redis,
		logger: logger,
	}
	return module
}

var ProviderSet = wire.NewSet(NewModule, wire.FieldsOf(new(*application.Application), "Router", "Database", "Redis", "Logger"))
