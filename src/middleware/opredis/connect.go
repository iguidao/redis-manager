package opredis

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"

	"github.com/go-redis/redis/v9"
)

type ClientConnect struct {
	*redis.Client
}

var RD ClientConnect

func ConnectRedis(addr string) bool {
	rd := redis.NewClient(&redis.Options{
		Addr: addr,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})
	RD = ClientConnect{rd}
	_, err := RD.Ping(ctx).Result()
	if err != nil {
		logger.Error("Redis Connect Error: ", err)
		return false
	}
	return true
}
