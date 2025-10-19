package middleware

import (
	"errors"
	"github.com/iqbalatma/gofortify"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string = c.GetHeader("Authorization")
		accessTokenVerifier, _ := c.Cookie("access_token_verifier")
		payload, err := gofortify.ValidateAccessToken(
			&token,
			&accessTokenVerifier,
		)

		if err != nil {
			var httpErr *utils.HTTPError

			switch err {
			case gofortify.ErrInvalidTokenType:
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
				err = gofortify.ErrJWTSubjectNotFound
			}
			err = errors.New("cannot find user")
		}
		c.Set("user", &user)
	}
}
