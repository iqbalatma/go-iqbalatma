package middleware

import (
	"errors"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if ginErr := c.Errors.Last(); ginErr != nil {
			originalErr := ginErr.Err
			config.AppLogger.WithFields(logrus.Fields{
				"request_id": c.GetString("RequestID"),
				"method":     c.Request.Method,
				"path":       c.Request.URL.Path,
				"ip":         c.ClientIP(),
			}).Error(originalErr)
			var httpError *exception.HTTPError
			if errors.As(originalErr, &httpError) {
				c.AbortWithStatusJSON(httpError.StatusCode, exception.HTTPError{
					Code:      httpError.Code,
					Message:   httpError.Message,
					Timestamp: httpError.Timestamp,
				})
				return
			}

			if errors.Is(originalErr, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, exception.NewHttpError(
					enum.ERR_NOT_FOUND,
					"Data not found",
					http.StatusNotFound,
				))
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, exception.HTTPError{
				Code:      enum.ERR_INTERNAL_SERVER_ERROR,
				Message:   originalErr.Error(),
				Timestamp: time.Now(),
			})
		}
	}
}
