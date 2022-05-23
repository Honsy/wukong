package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	ginLogger := NewDefaultLogrusLogger().WithPrefix("Gin")

	return func(ctx *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		ctx.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := ctx.Request.URL.Path

		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		ginLogger.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}
