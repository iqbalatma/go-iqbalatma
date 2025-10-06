package iqbalatma_go_jwt_authentication

import (
	"fmt"
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/blacklist"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAccessToken(jwtToken string, accessTokenVerifier *string, blacklist blacklist.Blacklist) (*Payload, error) {
	payload, err := Decode(jwtToken)
	if err != nil {
		return nil, err
	}

	// check token type, make sure this is access token
	if payload.TYPE != ACCESS_TOKEN {
		return nil, ErrInvalidTokenType
	}

	//check is on blacklist
	jti := blacklist.Get(payload.JTI)

	//when jti is on blacklist
	if jti != nil {
		return nil, ErrExpiredToken
	}

	//if now greater than exp, mean it's already expired
	if time.Now().Unix() > payload.EXP {
		return nil, ErrExpiredToken
	}

	//check is atv is valid
	if payload.IUC {
		if accessTokenVerifier == nil {
			return nil, ErrMissingRequiredAccessTokenVerifierCookie
		}

		fmt.Println("GET ATV FROM PAYLOAD HASHED : " + payload.ATV)
		fmt.Println("GET ATV UUID FROM COOKIE PLAIN  : " + *accessTokenVerifier)

		err := bcrypt.CompareHashAndPassword([]byte(payload.ATV), []byte(*accessTokenVerifier))
		if err != nil {
			return nil, ErrInvalidAccessTokenVerifier
		}
	}

	return payload, nil
}

func GetRemovedBearer(token string) string {
	token = strings.TrimSpace(token)
	splitToken := strings.SplitN(token, " ", 2)
	if len(splitToken) == 2 && splitToken[0] == "Bearer" {
		return splitToken[1]
	}

	return token
}
