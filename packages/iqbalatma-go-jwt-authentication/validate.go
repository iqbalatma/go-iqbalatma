package iqbalatma_go_jwt_authentication

import (
	"errors"
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/blacklist"
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/config"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAccessToken(jwtToken string, accessTokenVerifier *string) (*Payload, error) {
	payload, err := Decode(jwtToken)

	if err != nil {
		return nil, err
	}

	incidentTime, err := GetIncidentTime()

	//if not nil, it's mean incident time is not set, could be redis is broken, blacklist all jwt before this incident
	if err != nil {
		return nil, ErrExpiredToken
	}

	//it's mean this token is created before incident time, could be it's actually on blacklist but the list is gone
	//so blacklist all token that created before incident time
	if payload.IAT < incidentTime {
		return nil, ErrExpiredToken
	}

	// check token type, make sure this is access token
	if payload.TYPE != ACCESS_TOKEN {
		return nil, ErrInvalidTokenType
	}

	//check is on blacklist
	jti := blacklist.GetBlacklist().Get(payload.JTI)

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

func GetIncidentTime() (int64, error) {
	incidentTime := blacklist.GetBlacklist().Get(config.Config.IncidentKey)
	now := time.Now().Unix()

	if incidentTime == nil { //it's mean incident time is not set
		blacklist.GetBlacklist().Set(config.Config.IncidentKey, now, 0)
		return now, errors.New("incident time not set")
	}
	incidentTimeUnix, ok := incidentTime.(int64)
	if !ok {
		incidentTimeUnixString, ok := incidentTime.(string)
		if !ok {
			blacklist.GetBlacklist().Delete(config.Config.IncidentKey)
			blacklist.GetBlacklist().Set(config.Config.IncidentKey, now, 0)
			return now, errors.New("incident time is not int64")
		}

		incidentTimeUnix, err := strconv.ParseInt(incidentTimeUnixString, 10, 64)
		if err != nil {
			blacklist.GetBlacklist().Delete(config.Config.IncidentKey)
			blacklist.GetBlacklist().Set(config.Config.IncidentKey, now, 0)
			return now, errors.New("failed to parse incident time")
		}

		return incidentTimeUnix, nil
	}

	return incidentTimeUnix, nil
}
