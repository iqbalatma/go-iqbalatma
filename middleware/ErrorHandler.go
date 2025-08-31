package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iqbalatma/go-iqbalatma/app/enum"
	error2 "iqbalatma/go-iqbalatma/error"
	"net/http"
	"time"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, ginErr := range c.Errors {
			originalErr := ginErr.Err

			var httpErr *error2.HTTPError
			if errors.As(originalErr, &httpErr) {
				c.AbortWithStatusJSON(httpErr.StatusCode, map[string]string{
					"code":      string(httpErr.Code),
					"message":   httpErr.Message,
					"timestamp": httpErr.Timestamp.Format("2006-01-02 15:04:05"),
				})
				return
			}

			if errors.Is(originalErr, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
					"code":      string(enum.ERR_NOT_FOUND),
					"message":   "record not found",
					"timestamp": time.Now().Format("2006-01-02 15:04:05"),
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"code":      string(enum.ERR_INTERNAL_SERVER_ERROR),
				"message":   originalErr.Error(),
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
			return
		}
	}
}
