package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appStartTime := time.Now()

		path := ctx.FullPath()
		method := ctx.Request.Method
		fmt.Println("Path:", path)
		fmt.Println("Method:", method)

		ctx.Next()

		fmt.Println("Status:", ctx.Writer.Status())

		appExecutionTime := time.Since(appStartTime)
		fmt.Println(appExecutionTime)
	}
}
