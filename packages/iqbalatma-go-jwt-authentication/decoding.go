package iqbalatma_go_jwt_authentication

import (
	"github.com/golang-jwt/jwt/v5"
)

func Decoding(jwtString string) (*Payload, error) {
	key := []byte(Config.JWTSecretKey)
	payload := &Payload{}

	token, err := jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	//from parse to jwt claims
	claims, _ := token.Claims.(jwt.MapClaims)
	err = payload.FromMapClaims(claims)

	if err != nil {
		return nil, err
	}
	return payload, nil
}
