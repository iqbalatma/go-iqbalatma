package iqbalatma_go_jwt_authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getDefaultPayload() *Payload {
	return &Payload{
		ATV:  "s",
		ISS:  "localhost",
		IAT:  time.Now().Unix(),
		EXP:  time.Now().Unix(),
		NBF:  time.Now().Unix(),
		JTI:  uuid.New().String(),
		SUB:  "sub",
		IUA:  "iqbalatma-go-jwt-authentication",
		IUC:  true,
		TYPE: ACCESS_TOKEN,
	}
}

func Encode() (string, error) {
	key := []byte(Config.JWTSecretKey)

	payload := getDefaultPayload()
	payload.SUB = "IQBALATMA"
	payload.EXP = payload.EXP + 1000

	token := jwt.NewWithClaims(GetSigningMethod(),
		payload.ToMapClaims(),
	)

	return token.SignedString(key)
}
