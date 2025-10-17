package middleware

import (
	"errors"
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/config"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var httpResponse *utils.HTTPError

		if ginErr := c.Errors.Last(); ginErr != nil {
			originalErr := ginErr.Err

			logError(c, originalErr)
			fmt.Printf("Error type: %T\n", originalErr)
			if errors.Is(originalErr, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, utils.NewHttpError("Data not found", enum.ERR_NOT_FOUND))
				return
			}

			if errors.As(originalErr, &httpResponse) {
				c.AbortWithStatusJSON(httpResponse.StatusCode, utils.NewHttpError(httpResponse.Message, httpResponse.Code))
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewHttpError(originalErr.Error(), enum.ERR_INTERNAL_SERVER_ERROR))
		}
	}
}

func logError(c *gin.Context, err error) {
	config.AppLogger.WithFields(logrus.Fields{
		"request_id": c.GetString("RequestID"),
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"ip":         c.ClientIP(),
	}).Error(err)
}
