package opredis

import (
	"strings"

	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func CTypeKey(keyname string) (string, bool) {
	ok, err := CRD.Type(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Set key: ", keyname, " Error: ", err)
		return "", false
	}
	return ok, true
}

// String key op
func CGetStringKey(keyname string) (string, bool) {
	val, err := CRD.Get(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Get key: ", keyname, " Error: ", err)
		return "", false
	}
	return val, true
}

func CGetListKey(keyname string) ([]string, bool) {
	lnum, err := CRD.LLen(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis LLEN key: ", keyname, " Error: ", err)
		return nil, false
	}
	// if lnum > 100 {
	// 	lnum = 100
	// }
	val, err := CRD.LRange(ctx, keyname, 0, lnum).Result()
	if err != nil {
		logger.Error("Redis LRANGE key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}
func CGetHashKey(keyname string) (map[string]string, bool) {
	val, err := CRD.HGetAll(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis HGETALL key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}
func CGetSetKey(keyname string) ([]string, bool) {
	val, err := CRD.SMembers(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis SMEMBERS key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}
func CGetZsetKey(keyname string) ([]string, bool) {
	znum, err := CRD.ZCard(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis ZCARD key: ", keyname, " Error: ", err)
		return nil, false
	}
	// if znum > 100 {
	// 	znum = 100
	// }
	val, err := CRD.ZRange(ctx, keyname, 0, znum).Result()
	if err != nil {
		logger.Error("Redis ZRANGE key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}
func CDelKey(keyname string) (int64, bool) {
	val, err := CRD.Del(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Del key: ", keyname, " Error: ", err)
		return -1, false
	}
	return val, true
}
func CGetClusterNode() []string {
	clusternode := CRD.ClusterNodes(ctx)
	nodeinfo := strings.Split(clusternode.Val(), "\n")
	return nodeinfo
}
