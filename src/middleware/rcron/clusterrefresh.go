package rcron

import "github.com/iguidao/redis-manager/src/middleware/logger"

func ClusterRefresh() {
	logger.Info("定时任务：刷新cluster任务启动")
}
