package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	//"goblog.com/pkg/application"
	"goblog.com/pkg/jwtauth"
	//"goblog.com/pkg/session"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module struct {
	router *gin.Engine
	database *gorm.DB
	redis redis.UniversalClient
	logger *zap.Logger
	session *session.Session
	validator *validator.Validate
	jwtAuth *jwtauth.JWTAuth

	webHandler *webHandler
	apiHandler *apiHandler
	service    *service
	repository *repository
}

var module *Module

func NewModule(
		router *gin.Engine, database *gorm.DB, redis redis.UniversalClient, logger *zap.Logger,
		session *session.Session, jwtAuth *jwtauth.JWTAuth, validator *validator.Validate,
		webHandler *webHandler, apiHandler *apiHandler, service *service, repository *repository,
	) *Module {
	module = &Module{
		router: router,
		database: database,
		redis: redis,
		logger: logger,
		session: session,
		validator: validator,
		jwtAuth: jwtAuth,

		webHandler: webHandler,
		apiHandler: apiHandler,
		service:    service,
		repository: repository,
	}
	return module
}

var ProviderSet = wire.NewSet(
		NewModule, NewHandler, NewApiHandler, NewService, NewRepository,
		wire.FieldsOf(new(*application.Application), "Logger", "Database", "Redis", "Router", "Session", "Validator", "JwtAuth"),
	)

// registry
func (m *Module) Registry() error {
	m.RegisterRouter()

	if err := m.AutoMigrate(); err != nil {
		return err
	}

	err := m.validator.RegisterValidation("is_exists", func(fl validator.FieldLevel) bool {
		user := m.repository.firstByUsername(fl.Field().String())
		return user == nil
	})
	if err != nil {
		return errors.New("validator register error")
	}

	return nil
}

// migrate
func (m *Module) AutoMigrate() error {
	return m.database.AutoMigrate(&User{});
}