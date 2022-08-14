package opredis

import (
	"context"
	"strconv"
	"time"

	"github.com/iguidao/redis-manager/src/middleware/logger"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

// Base op
func TypeKey(keyname string) (string, bool) {
	ok, err := RD.Type(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Set key: ", keyname, " Error: ", err)
		return "", false
	}
	return ok, true
}

func TtlKey(keyname string) (time.Duration, bool) {
	val, err := RD.TTL(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Get key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}

func GetSlowLog() ([]redis.SlowLog, bool) {
	val, err := RD.SlowLogGet(ctx, 100).Result()
	if err != nil {
		logger.Error("Redis Get Slowlog Error: ", err)
		return nil, false
	}
	return val, true
}

func GetScanKey(cursor uint64, allnum int64) ([]string, uint64, bool) {
	keys, val, err := RD.Scan(ctx, cursor, "*", allnum).Result()
	if err != nil {
		logger.Error("Redis Get Scan "+strconv.FormatUint(cursor, 10)+"Error: ", err)
		return nil, 0, false
	}
	return keys, val, true
}

func BgsaveKey() (string, bool) {

	val, err := RD.BgSave(ctx).Result()
	if err != nil {
		logger.Error("Redis Bgsave Error: ", err)
		return "Execution failed", false
	}
	return val, true
}

func DebugKey(keyname string) (string, bool) {

	val, err := RD.DebugObject(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Debug Key:", keyname, "  Error: ", err)
		return "Execution Debug failed", false
	}
	return val, true
}

// String key op
func GetStringKey(keyname string) (string, bool) {
	val, err := RD.Get(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Get key: ", keyname, " Error: ", err)
		return "", false
	}
	return val, true
}

func SizeStringKey(keyname string) (int64, bool) {
	val, err := RD.StrLen(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Strlen key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}

// list key op
func GetListKey(keyname string) ([]string, bool) {
	lnum, err := RD.LLen(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis LLEN key: ", keyname, " Error: ", err)
		return nil, false
	}
	if lnum > 100 {
		lnum = 100
	}
	val, err := RD.LRange(ctx, keyname, 0, lnum).Result()
	if err != nil {
		logger.Error("Redis LRANGE key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}

func SizeListKey(keyname string) (int64, bool) {
	val, err := RD.LLen(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis LLEN key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}

// Hash key op
func GetHashKey(keyname string) (map[string]string, bool) {
	val, err := RD.HGetAll(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis HGETALL key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}

func SizeHashKey(keyname string) (int64, bool) {
	val, err := RD.HLen(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis Hlen key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}

//Set key op
func GetSetKey(keyname string) ([]string, bool) {
	val, err := RD.SMembers(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis SMEMBERS key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}
func SizeSetKey(keyname string) (int64, bool) {
	val, err := RD.SCard(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis SMEMBERS key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}

//Zset key op
func GetZsetKey(keyname string) ([]string, bool) {
	znum, err := RD.ZCard(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis ZCARD key: ", keyname, " Error: ", err)
		return nil, false
	}
	if znum > 100 {
		znum = 100
	}
	val, err := RD.ZRange(ctx, keyname, 0, znum).Result()
	if err != nil {
		logger.Error("Redis ZRANGE key: ", keyname, " Error: ", err)
		return nil, false
	}
	return val, true
}

func SizeZsetKey(keyname string) (int64, bool) {
	val, err := RD.ZCard(ctx, keyname).Result()
	if err != nil {
		logger.Error("Redis ZCARD key: ", keyname, " Error: ", err)
		return 0, false
	}
	return val, true
}
