package mysql

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
)

// cloud redis list
func (m *MySQL) GetAllCloudredis() []CloudInfo {
	var cloudinfo []CloudInfo
	m.Find(&cloudinfo)
	return cloudinfo
}
func (m *MySQL) GetCloudNumber(cloud string) int64 {
	var cloudinfo []CloudInfo
	var count int64
	m.Model(cloudinfo).Where("cloud = ?", cloud).Find(&cloudinfo).Count(&count)
	return count
}
func (m *MySQL) GetCloudredis(cloud, region string) []CloudInfo {
	var cloudinfo []CloudInfo
	m.Model(cloudinfo).Where("cloud = ? AND region = ?", cloud, region).Find(&cloudinfo)
	return cloudinfo
}
func (m *MySQL) GetCloudRegion() []CloudInfo {
	var cloudinfo []CloudInfo
	m.Select([]string{"cloud", "region"}).Find(&cloudinfo)
	return cloudinfo
}
func (m *MySQL) ExistCloudredisId(cloud, instanceId string) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(cloudinfo).Where("cloud = ? AND instance_id = ?", cloud, instanceId).First(&cloudinfo).Error; err != nil {
		return false
	}
	return true
}
func (m *MySQL) UpdateCloudPassword(cloud, instanceid, password string) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(cloudinfo).Where("cloud = ? AND instance_id = ?", cloud, instanceid).Update("password", password).Error; err != nil {
		logger.Error("update cloud password error: ", err)
		return false
	}
	return true
}
func (m *MySQL) GetCloudAddress(cloud, instanceid string) (string, string, int) {
	var cloudinfo *CloudInfo
	m.Where("instance_id = ? AND cloud = ?", instanceid, cloud).First(&cloudinfo)
	return cloudinfo.Password, cloudinfo.PrivateIp, cloudinfo.Port
}

// tx
func (m *MySQL) UppdateTxCloudRedis(cloud string, redisinfo model.TxLResponseInstanceSet) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(&cloudinfo).Where("cloud = ? AND instance_id = ?", cloud, redisinfo.InstanceId).Updates(map[string]interface{}{
		"instance_name":      redisinfo.InstanceName,
		"private_ip":         redisinfo.WanIp,
		"port":               redisinfo.Port,
		"region":             redisinfo.Region,
		"createtime":         redisinfo.Createtime,
		"size":               redisinfo.Size,
		"instance_status":    redisinfo.InstanceTitle,
		"redis_shard_size":   redisinfo.RedisShardSize,
		"redis_shard_num":    redisinfo.RedisShardNum,
		"redis_replicas_num": redisinfo.RedisReplicasNum,
		"no_auth":            redisinfo.NoAuth,
		"public_ip":          redisinfo.WanAddress,
	}).Error; err != nil {
		logger.Error("update cloud instanceid: ", redisinfo.InstanceId, "info error: ", err)
		return false
	}
	return true
}

func (m *MySQL) AddTxCloudRedis(cloud string, redisinfo model.TxLResponseInstanceSet) (int, bool) {
	addcluster := &CloudInfo{
		Cloud:            cloud,
		InstanceId:       redisinfo.InstanceId,
		InstanceName:     redisinfo.InstanceName,
		PrivateIp:        redisinfo.WanIp,
		Port:             redisinfo.Port,
		Region:           redisinfo.Region,
		Createtime:       redisinfo.Createtime,
		Size:             redisinfo.Size,
		InstanceStatus:   redisinfo.InstanceTitle,
		RedisShardSize:   redisinfo.RedisShardSize,
		RedisShardNum:    redisinfo.RedisShardNum,
		RedisReplicasNum: redisinfo.RedisReplicasNum,
		NoAuth:           redisinfo.NoAuth,
		PublicIp:         redisinfo.WanAddress,
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		logger.Error("add cloud redis error: ", result.Error)
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}

// ali
func (m *MySQL) UppdateAliCloudRedis(cloud string, redisinfo model.AliRedisInstancesKVStoreInstance) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(&cloudinfo).Where("cloud = ? AND instance_id = ?", cloud, redisinfo.InstanceId).Updates(map[string]interface{}{
		"instance_name":   redisinfo.InstanceName,
		"private_ip":      redisinfo.ConnectionDomain,
		"port":            redisinfo.Port,
		"region":          redisinfo.RegionId,
		"createtime":      redisinfo.CreateTime,
		"size":            redisinfo.Capacity,
		"instance_status": redisinfo.InstanceStatus,
	}).Error; err != nil {
		logger.Error("update cloud instanceid: ", redisinfo.InstanceId, "info error: ", err)
		return false
	}
	return true
}

func (m *MySQL) AddAliCloudRedis(cloud string, redisinfo model.AliRedisInstancesKVStoreInstance) (int, bool) {
	addcluster := &CloudInfo{
		Cloud:          cloud,
		InstanceId:     redisinfo.InstanceId,
		InstanceName:   redisinfo.InstanceName,
		PrivateIp:      redisinfo.ConnectionDomain,
		Port:           redisinfo.Port,
		Region:         redisinfo.RegionId,
		Createtime:     redisinfo.CreateTime,
		Size:           redisinfo.Capacity,
		InstanceStatus: redisinfo.InstanceStatus,
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		logger.Error("add cloud redis error: ", result.Error)
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}
func (m *MySQL) DelCloud(instanceid string) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(cloudinfo).Where("instance_id = ?", instanceid).Delete(&cloudinfo).Error; err != nil {
		logger.Error(err)
		return false
	}
	return true
}
