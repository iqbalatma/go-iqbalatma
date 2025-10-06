package blacklist

import "time"

type Blacklist interface {
	Get(key string) any
	Set(key string, value any, expired time.Duration)
	Delete(key string)
}
