package opredis

import (
	"github.com/go-redis/redis/v9"
)

func SlowKey() []redis.SlowLog {
	val, slowlogok := GetSlowLog()
	if !slowlogok {
		return nil
	}
	return val
}
