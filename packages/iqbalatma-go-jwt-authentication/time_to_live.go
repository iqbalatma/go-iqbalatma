package iqbalatma_go_jwt_authentication

func GetAccessTTL() int {
	return Config.AccessTokenTTL
}

func GetRefreshTTL() int {
	return Config.RefreshTokenTTL
}
