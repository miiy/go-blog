package health

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	ErrDatabase = errors.New("database: connection error")
	ErrRedis = errors.New("redis: connection error")
)

const (
	SqlShowTable = "SHOW TABLES"
)

func livenessHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}

func readinessHandler(ctx *gin.Context) {
	var errMsg []string

	var rst []string
    if err := module.db.Raw(SqlShowTable).Scan(&rst).Error; err != nil {
		errMsg = append(errMsg, ErrDatabase.Error())
	}

	c := context.Background()
	if err := module.redis.Ping(c).Err(); err != nil {
		errMsg = append(errMsg, ErrRedis.Error())
	}

	if len(errMsg) == 0 {
		ctx.String(http.StatusOK, "Ok")
	} else {
		ctx.String(http.StatusInternalServerError, strings.Join(errMsg, "\n"))
	}
	return
}
