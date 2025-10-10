package iqbalatma_go_jwt_authentication

import (
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/blacklist"
	"time"
)

func Revoke(jwtToken string) (*Payload, error) {
	payload, err := Decode(jwtToken)
	if err != nil {
		return nil, err
	}
	ttl := time.Unix(payload.EXP, 0).Sub(time.Now())
	blacklist.AddBlacklistToken(payload.JTI, ttl)

	return payload, nil
}
