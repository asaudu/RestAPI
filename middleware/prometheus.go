package middleware

import (
	"strconv"
	"time"

	"addyCodes.com/RestAPI/metrics"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.FullPath()

		ctx.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa((ctx.Writer.Status()))

		metrics.HttpRequestTotal.WithLabelValues(path, c.Request.Method, status).Inc()
		metrics.HttpRequestDuration.WithLabelValues(path).Observe(duration)
	}
}
