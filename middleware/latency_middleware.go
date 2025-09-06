package middleware

import (
	"iqbalatma/go-iqbalatma/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLatencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		config.AppLogger.WithFields(logrus.Fields{"status": c.Writer.Status(),
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"client_ip": c.ClientIP(),
			"latency":   latency.Milliseconds(),
		}).Info("request completed")
	}
}
