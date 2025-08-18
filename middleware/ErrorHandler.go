package middleware

import (
	"github.com/gin-gonic/gin"
	error2 "iqbalatma/go-iqbalatma/error"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *error2.HTTPError:
				c.AbortWithStatusJSON(e.StatusCode, map[string]string{
					"code":      string(e.Code),
					"message":   e.Message,
					"timestamp": e.Timestamp.Format("2006-01-02 15:04:05"),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}
