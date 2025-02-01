package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func RequestLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		ip := c.ClientIP()

		c.Next()

		status := c.Writer.Status()
		latency := time.Since(start)

		logger.Infof("[REQUEST] %s %s | Status: %d | IP: %s | Latency: %v", method, path, status, ip, latency)
	}
}
