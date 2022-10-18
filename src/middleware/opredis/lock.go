package opredis

import (
	"time"

	"github.com/iguidao/redis-manager/src/cfg"
)

func LockCheck(key string, keytime time.Duration) bool {
	if ConnectRedis(cfg.Get_Info("REDIS")) {
		if LockOp("Lock-"+key, keytime) {
			return true
		}
	}
	return false
}

func LockRm(key string) bool {
	if ConnectRedis(cfg.Get_Info("REDIS")) {
		if UnLockOp("Lock-" + key) {
			return true
		}
	}
	return false
}
