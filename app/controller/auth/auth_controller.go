package auth

import (
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/service/auth"
	iqbalatma_go_jwt_authentication "iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication"
	"iqbalatma/go-iqbalatma/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func (this AuthController) Authenticate(c *gin.Context) error {
	user, err := this.AuthService.Authenticate(c)
	if err != nil {
		return err
	}

	accessToken, atv, err := iqbalatma_go_jwt_authentication.Encode(user,
		iqbalatma_go_jwt_authentication.ACCESS_TOKEN,
		true,
		"localhost",
		c.Request.UserAgent(),
	)

	refreshToken, _, err := iqbalatma_go_jwt_authentication.Encode(user,
		iqbalatma_go_jwt_authentication.REFRESH_TOKEN,
		true,
		"localhost",
		c.Request.UserAgent(),
	)
	if err != nil {
		return err
	}

	c.SetCookie(
		"refresh_token",
		refreshToken,
		3600,
		"/",
		"localhost",
		true,
		true,
	)

	c.SetCookie(
		"access_token_verifier",
		atv,
		3600,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(http.StatusOK, &utils.HTTPResponse{
		Message:   "Authenticate User Successfully",
		Timestamp: time.Now(),
		Code:      enum.SUCCESS,
		Payload: &utils.Payload{
			Data: map[string]interface{}{
				"id":           user.ID.String(),
				"email":        user.Email,
				"first_name":   user.FirstName,
				"last_name":    user.LastName,
				"access_token": accessToken,
			},
		},
	})
	return nil
}

func (this AuthController) Logout(c *gin.Context) error {
	var accessToken string = c.GetHeader("Authorization")
	_, err := iqbalatma_go_jwt_authentication.Revoke(
		iqbalatma_go_jwt_authentication.GetRemovedBearer(accessToken),
	)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, &utils.HTTPResponse{
		Message:   "Logout Successfully",
		Timestamp: time.Now(),
		Code:      enum.SUCCESS,
	})
	return nil
}

func NewAuthController() AuthController {
	return AuthController{
		AuthService: auth.NewAuthService(),
	}
}
