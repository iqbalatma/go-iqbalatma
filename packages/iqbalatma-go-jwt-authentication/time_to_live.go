package iqbalatma_go_jwt_authentication

import "iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/config"

func GetAccessTTL() int {
	return config.Config.AccessTokenTTL
}

func GetRefreshTTL() int {
	return config.Config.RefreshTokenTTL
}
