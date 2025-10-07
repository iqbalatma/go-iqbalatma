package iqbalatma_go_jwt_authentication

import (
	"fmt"
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func getDefaultPayload() *Payload {
	return &Payload{
		ATV:  "",
		ISS:  "",
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

func Encode(
	subject JWTSubject,
	tokenType TokenType,
	iuc bool,
	iss string,
	iua string,
) (string, string, error) {
	key := []byte(config.Config.JWTSecretKey)

	payload := getDefaultPayload()
	incidentTime, err := GetIncidentTime()
	if err != nil { //it's mean incident time is not set or wrong form, then it will set at this time
		payload.EXP = incidentTime
		payload.NBF = incidentTime
		payload.IAT = incidentTime
	}
	payload.TYPE = tokenType
	payload.SUB = subject.GetSubjectKey()
	payload.ISS = iss
	payload.IUA = iua
	payload.IUC = iuc

	atv, err := addATV(payload)
	if err != nil {
		return "", "", err
	}

	addTTL(payload)

	token := jwt.NewWithClaims(config.GetSigningMethod(),
		payload.ToMapClaims(),
	)

	signedString, err := token.SignedString(key)
	if err != nil {
		return "", "", err
	}
	return signedString, atv, err
}

func addTTL(payload *Payload) {
	if payload.TYPE == ACCESS_TOKEN {
		payload.EXP = time.Now().Add(time.Duration(GetAccessTTL()) * time.Minute).Unix()
	} else {
		payload.EXP = time.Now().Add(time.Duration(GetRefreshTTL()) * time.Minute).Unix()
	}
}

func addATV(payload *Payload) (string, error) {
	if payload.TYPE == ACCESS_TOKEN {
		var atv string = uuid.New().String()
		fmt.Println("ATV UUID CREATED : " + atv)
		bytes, err := bcrypt.GenerateFromPassword([]byte(atv), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		fmt.Println("ATV UUID HASHED : " + string(bytes))
		payload.ATV = string(bytes)

		return atv, nil
	}

	return "", nil
}
