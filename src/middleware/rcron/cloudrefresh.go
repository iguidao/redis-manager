package rcron

import (
	"encoding/json"
	"log"

	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/txcloud"
	"github.com/iguidao/redis-manager/src/middleware/util"
)

func CloudRefresh() {
	log.Println("qidong")
	logger.Info("定时任务：刷新云redis任务启动")
	cloudset := make(map[string]string)
	cloudinfo := mysql.DB.GetCloudRegion()
	if len(cloudinfo) == 0 {
		logger.Info("定时任务：数据库没有数据，不用获取最新的云redis数据")
		return
	}
	for _, v := range cloudinfo {
		cloudset[v.Region] = v.Cloud
	}
	for i, v := range cloudset {
		if v == "txredis" {
			if !txcloud.TxRedisContent(i) {
				logger.Error("定时任务：链接腾讯云redis失败")
				return
			} else {
				list, ok := txcloud.TxListRedis()
				var rlist model.TxL
				if ok {
					err := json.Unmarshal([]byte(list), &rlist)
					if err == nil {
						go util.TxWriteRedis(v, rlist)
						logger.Info("定时任务：开始更新腾讯云redis数据")
					} else {
						logger.Error("定时任务：json解析云redis数据失败", err)
					}
				} else {
					logger.Error("定时任务：获取云redis数据失败")
				}
			}
		}
	}

}
