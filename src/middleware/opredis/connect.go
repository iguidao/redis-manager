package opredis

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"

	"github.com/go-redis/redis/v9"
)

// 单点和codis链接
type ClientConnect struct {
	*redis.Client
}

var RD ClientConnect

func ConnectRedis(addr, password string) bool {
	rd := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		// DB:       0,        // use default DB
	})
	RD = ClientConnect{rd}
	_, err := RD.Ping(ctx).Result()
	if err != nil {
		logger.Error("Redis Connect Error: ", err)
		return false
	}
	return true
}

// 集群链接
type ClientClusterConnect struct {
	*redis.ClusterClient
}

var CRD ClientClusterConnect

func ConnectRedisCluster(addr []string, password string) bool {
	rd := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addr,
		Password: password,
		// DialTimeout:  200 * time.Microsecond,
		// ReadTimeout:  200 * time.Microsecond,
		// WriteTimeout: 200 * time.Microsecond,
	})
	CRD = ClientClusterConnect{rd}
	return true
}
