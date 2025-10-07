package iqbalatma_go_jwt_authentication

import (
	"fmt"
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/blacklist"
	"time"
)

func Revoke(jwtToken string) (*Payload, error) {
	fmt.Println("Revoke called " + jwtToken)
	payload, err := Decode(jwtToken)
	if err != nil {
		return nil, err
	}
	ttl := time.Unix(payload.EXP, 0).Sub(time.Now())
	blacklist.AddBlacklistToken(payload.JTI, ttl)

	return payload, nil
}
