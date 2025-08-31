package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"iqbalatma/go-iqbalatma/config"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		config.AppLogger.WithField("Request ID", requestID)
		c.Next()
	}
}
