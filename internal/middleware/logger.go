package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		status := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"status":   status,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"latency":  latency.String(),
			"clientIP": c.ClientIP(),
		}).Info("HTTP request processed")
	}
}
