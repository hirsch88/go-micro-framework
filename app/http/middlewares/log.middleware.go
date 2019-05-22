package middlewares

import (
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
)

type LogMiddleware struct {
	log *zap.SugaredLogger
}

func (m *LogMiddleware) Handler(ctx *gin.Context) {
	// before request
	t := time.Now()

	ctx.Next()

	// after request
	latency := time.Since(t)
	m.log.Info(latency)

	// access the status we are sending
	status := ctx.Writer.Status()
	m.log.Info(status)
}

func NewLogMiddleware(log *zap.SugaredLogger) *LogMiddleware {
	return &LogMiddleware{log}
}
