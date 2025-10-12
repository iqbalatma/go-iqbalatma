package auth

import (
	"fmt"
	"github.com/iqbalatma/gofortress"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/app/interface/service"
	"iqbalatma/go-iqbalatma/app/model"
	"iqbalatma/go-iqbalatma/app/service/auth"
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

	accessToken, atv, err := gofortress.Encode(user,
		gofortress.ACCESS_TOKEN,
		true,
		"localhost",
		c.Request.UserAgent(),
	)

	refreshToken, _, err := gofortress.Encode(user,
		gofortress.REFRESH_TOKEN,
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
	_, err := gofortress.Revoke(
		gofortress.GetRemovedBearer(accessToken),
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

func (this AuthController) Refresh(c *gin.Context) error {
	refreshToken, _ := c.Cookie("refresh_token")

	_, err := gofortress.Revoke(
		gofortress.GetRemovedBearer(refreshToken),
	)

	if err != nil {
		return err
	}

	value, exists := c.Get("user")
	fmt.Println(value)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return nil
	}

	user, ok := value.(*model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cast user"})
		return nil
	}

	accessToken, atv, err := gofortress.Encode(user,
		gofortress.ACCESS_TOKEN,
		true,
		"localhost",
		c.Request.UserAgent(),
	)

	refreshToken, _, err = gofortress.Encode(user,
		gofortress.REFRESH_TOKEN,
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
		Message:   "Refresh Token User Successfully",
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

func NewAuthController() AuthController {
	return AuthController{
		AuthService: auth.NewAuthService(),
	}
}
