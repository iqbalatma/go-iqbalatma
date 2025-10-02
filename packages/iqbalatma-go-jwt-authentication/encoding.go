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
		SUB:  "",
		IUA:  "",
		IUC:  true,
		TYPE: ACCESS_TOKEN,
	}
}

func Encode(subject JWTSubject[any], tokenType TokenType) (string, error) {
	key := []byte(Config.JWTSecretKey)

	payload := getDefaultPayload()
	payload.TYPE = tokenType
	payload.SUB = subject.GetSubjectKey()
	addTTL(payload)

	token := jwt.NewWithClaims(GetSigningMethod(),
		payload.ToMapClaims(),
	)

	return token.SignedString(key)
}

func addTTL(payload *Payload) {
	if payload.TYPE == ACCESS_TOKEN {
		payload.EXP = time.Now().Add(time.Duration(GetAccessTTL()) * time.Minute).Unix()
	} else {
		payload.EXP = time.Now().Add(time.Duration(GetRefreshTTL()) * time.Minute).Unix()
	}
}
