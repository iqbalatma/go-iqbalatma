package middleware

import (
	"errors"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/iqbalatma/gofortress"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RefreshMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("refresh_token")
		payload, err := gofortress.ValidateRefreshToken(
			&token,
		)

		if err != nil {
			var httpErr *utils.HTTPError

			switch err {
			case gofortress.ErrInvalidTokenType:
				httpErr = exception.InvalidTokenTypeException()
			}

			if httpErr == nil {
				httpErr = exception.UnauthorizedException(err.Error())
			}

			c.AbortWithStatusJSON(httpErr.StatusCode, utils.NewHttpError(httpErr.Message, httpErr.Code))
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
