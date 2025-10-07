package middleware

import (
	"errors"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/config"
	exception "iqbalatma/go-iqbalatma/error"
	iqbalatma_go_jwt_authentication "iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication"
	"iqbalatma/go-iqbalatma/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string = c.GetHeader("Authorization")
		accessTokenVerifier, _ := c.Cookie("access_token_verifier")
		payload, err := iqbalatma_go_jwt_authentication.ValidateAccessToken(
			iqbalatma_go_jwt_authentication.GetRemovedBearer(token),
			&accessTokenVerifier,
		)

		if err != nil {
			var httpErr *exception.HTTPError

			switch err {
			case iqbalatma_go_jwt_authentication.ErrInvalidTokenType:
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
				err = iqbalatma_go_jwt_authentication.ErrJWTSubjectNotFound
			}
			err = errors.New("cannot find user")
		}
		c.Set("user", &user)
	}
}
