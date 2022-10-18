package opredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func RedisSave(serverip string) bool {
	knowtime := time.Now().Unix()
	c, err := redis.Dial("tcp", serverip)
	if err != nil {
		logger.Debug("conn redis failed,", err)
		return false
	}
	defer c.Close()

	_, err = redis.String(c.Do("KSBGSAVE"))
	if err != nil {
		logger.Debug("ip: "+serverip+" 执行redis的KSBGSAVE操作失败：", err)
		_, err = redis.String(c.Do("BGSAVE"))
		if err != nil {
			logger.Debug("ip: "+serverip+" 执行redis的BGSAVE操作失败：", err)
			return false
		}
		go savebigkey(knowtime, serverip)
	}
	go savebigkey(knowtime, serverip)
	return true
}

func savebigkey(cmdtime int64, serverip string) bool {
	c, err := redis.Dial("tcp", serverip)
	if err != nil {
		logger.Debug("conn redis failed,", err)
		return false
	}
	defer c.Close()
	fornum := 0
	for {
		knowtime, err := redis.Int64(c.Do("LASTSAVE"))
		if err != nil {
			logger.Debug("ip: "+serverip+" 执行redis的LASTSAVE操作失败：", err)
			break
		}
		if knowtime > cmdtime {
			break
		}
		if fornum > 600 {
			logger.Error("ip: " + serverip + " 执行redis的LASTSAVE次数超过10min")
			break
		}
		fornum = fornum + 1
		time.Sleep(time.Duration(1) * time.Second)
	}
	return true
}
