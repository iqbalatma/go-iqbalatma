package iqbalatma_go_jwt_authentication

import (
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SigningMethod   string
	JWTSecretKey    string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

var Config *JWTConfig

func LoadJWTConfig() {
	accessTokenTTL, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_TTL"))
	if err != nil {
		accessTokenTTL = 30
	}

	refreshTokenTTL, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_TTL"))
	if err != nil {
		refreshTokenTTL = 10080
	}

	Config = &JWTConfig{
		SigningMethod:   os.Getenv("JWT_SIGNING_METHOD"),
		JWTSecretKey:    os.Getenv("JWT_SECRET_KEY"),
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
	}
}

func GetSigningMethod() jwt.SigningMethod {
	signingMethods := map[string]jwt.SigningMethod{
		"HS256": jwt.SigningMethodHS256,
		"HS384": jwt.SigningMethodHS384,
		"HS512": jwt.SigningMethodHS512,
		"ES512": jwt.SigningMethodES512,
		"ES384": jwt.SigningMethodES384,
		"ES256": jwt.SigningMethodES256,
		"EdDSA": jwt.SigningMethodEdDSA,
		"PS256": jwt.SigningMethodPS256,
		"PS512": jwt.SigningMethodPS512,
		"PS384": jwt.SigningMethodPS384,
		"RS256": jwt.SigningMethodRS256,
		"RS512": jwt.SigningMethodRS512,
		"RS384": jwt.SigningMethodRS384,
	}

	return signingMethods[Config.SigningMethod]
}
