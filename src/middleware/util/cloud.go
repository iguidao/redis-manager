package util

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func TxWriteRedis(cloud string, rlist model.TxL) {
	for _, v := range rlist.Response.InstanceSet {
		if !mysql.DB.ExistCloudredisId(v.InstanceId) {
			id, ok := mysql.DB.AddTxCloudRedis(cloud, v)
			if ok {
				logger.Info("write ", cloud, " redis to mysql ok: ", id, "instanceid: ", v.InstanceId)
			} else {
				logger.Error("write ", cloud, " redis to mysql false: ", id, "instanceid: ", v.InstanceId)
			}
		} else {
			ok := mysql.DB.UppdateTxCloudRedis(cloud, v)
			if ok {
				logger.Info("update ", cloud, " redis to mysql ok: ", "instanceid: ", v.InstanceId)
			} else {
				logger.Error("update ", cloud, " redis to mysql false: ", "instanceid: ", v.InstanceId)
			}
		}
	}
}

func AliWriteRedis(cloud string, rlist model.AliRedis) {
	for _, v := range rlist.Instances.KVStoreInstance {
		if !mysql.DB.ExistCloudredisId(v.InstanceId) {
			id, ok := mysql.DB.AddAliCloudRedis(cloud, v)
			if ok {
				logger.Info("write ", cloud, " redis to mysql ok: ", id, "instanceid: ", v.InstanceId)
			} else {
				logger.Error("write ", cloud, " redis to mysql false: ", id, "instanceid: ", v.InstanceId)
			}
		} else {
			ok := mysql.DB.UppdateAliCloudRedis(cloud, v)
			if ok {
				logger.Info("update ", cloud, " redis to mysql ok: ", "instanceid: ", v.InstanceId)
			} else {
				logger.Error("update ", cloud, " redis to mysql false: ", "instanceid: ", v.InstanceId)
			}
		}
	}
}
