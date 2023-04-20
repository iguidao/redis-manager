package util

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func TxWriteRedis(cloud string, rlist model.TxL) {
	for _, v := range rlist.Response.InstanceSet {
		if !mysql.DB.ExistCloudredisId(v.InstanceId) {
			id, ok := mysql.DB.AddCloudRedis(cloud, v)
			if ok {
				logger.Info("write ", cloud, " redis to mysql ok: ", id, "instanceid: ", v.InstanceId)
			} else {
				logger.Error("write ", cloud, " redis to mysql false: ", id, "instanceid: ", v.InstanceId)
			}
		}
	}
}
