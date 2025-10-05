package iqbalatma_go_jwt_authentication

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAccessToken(jwtToken string, accessTokenVerifier *string) (*Payload, error) {
	payload, err := Decode(jwtToken)
	if err != nil {
		return nil, err
	}

	// check token type, make sure this is access token
	if payload.TYPE != ACCESS_TOKEN {
		return nil, ErrInvalidTokenType
	}

	//check is on blacklist

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
