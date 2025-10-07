package blacklist

import (
	"iqbalatma/go-iqbalatma/packages/iqbalatma-go-jwt-authentication/config"
	"time"
)

type Blacklist interface {
	Get(key string) any
	Set(key string, value any, expired time.Duration)
	Delete(key string)
}

func GetBlacklist() Blacklist {
	return NewRedisBlacklist(config.RDB)
}
