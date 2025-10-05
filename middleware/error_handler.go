package middleware

import (
	"errors"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	"iqbalatma/go-iqbalatma/utils"
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
				c.AbortWithStatusJSON(httpError.StatusCode, exception.NewHttpError(
					httpError.Code,
					httpError.Message,
					httpError.StatusCode,
				))

				c.AbortWithStatusJSON(httpError.StatusCode, &utils.HTTPResponse{
					Message:   httpError.Message,
					Timestamp: httpError.Timestamp,
					Code:      httpError.Code,
				})
				return
			}

			if errors.Is(originalErr, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, &utils.HTTPResponse{
					Message:   "Data not found",
					Timestamp: time.Now(),
					Code:      enum.ERR_NOT_FOUND,
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, &utils.HTTPResponse{
				Message:   originalErr.Error(),
				Timestamp: time.Now(),
				Code:      enum.ERR_INTERNAL_SERVER_ERROR,
			})
		}
	}
}
