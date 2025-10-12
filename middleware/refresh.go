package middleware

import (
	"errors"
	"github.com/iqbalatma/gofortress"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("refresh_token")
		payload, err := gofortress.ValidateRefreshToken(
			gofortress.GetRemovedBearer(token),
		)

		if err != nil {
			var httpErr *exception.HTTPError

			switch err {
			case gofortress.ErrInvalidTokenType:
				httpErr = exception.InvalidTokenTypeException()
			}

			if httpErr == nil {
				httpErr = exception.UnauthorizedException(err.Error())
			}

			c.AbortWithStatusJSON(httpErr.StatusCode, &utils.HTTPResponse{
				Message:   httpErr.Message,
				Timestamp: httpErr.Timestamp,
				Code:      httpErr.Code,
			})
			return
		}

		var user model.User
		result := config.DB.Where("id = ?", payload.SUB).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				err = gofortress.ErrJWTSubjectNotFound
			}
			err = errors.New("cannot find user")
		}
		c.Set("user", &user)
	}
}
