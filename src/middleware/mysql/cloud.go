package mysql

import (
	"github.com/iguidao/redis-manager/src/middleware/model"
)

// cloud redis list
func (m *MySQL) GetAllCloudredis() []CloudInfo {
	var cloudinfo []CloudInfo
	m.Find(&cloudinfo)
	return cloudinfo
}
func (m *MySQL) GetCloudredis(cloud, region string) []CloudInfo {
	var cloudinfo []CloudInfo
	m.Find(&cloudinfo).Where("cloud = ? AND region = ?", cloud, region)
	return cloudinfo
}
func (m *MySQL) ExistCloudredisId(instanceId string) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(cloudinfo).Where("instance_id = ?", instanceId).First(&cloudinfo).Error; err != nil {
		return false
	}
	return true
}
func (m *MySQL) UpdateCloudPassword(instanceid, password string) bool {
	var cloudinfo *CloudInfo
	if err := m.Model(cloudinfo).Where("instance_id = ?", instanceid).Update("password", password).Error; err != nil {
		return false
	}
	return true
}
func (m *MySQL) GetCloudAddress(cloud, instanceid string) (string, string, int) {
	var cloudinfo *CloudInfo
	m.Where("instance_id = ? AND cloud = ?", instanceid, cloud).First(&cloudinfo)
	return cloudinfo.Password, cloudinfo.PrivateIp, cloudinfo.Port
}

func (m *MySQL) AddCloudRedis(cloud string, redisinfo model.TxLResponseInstanceSet) (int, bool) {
	addcluster := &CloudInfo{
		Cloud:          cloud,
		InstanceId:     redisinfo.InstanceId,
		InstanceName:   redisinfo.InstanceName,
		PrivateIp:      redisinfo.WanIp,
		Port:           redisinfo.Port,
		Region:         redisinfo.Region,
		Createtime:     redisinfo.Createtime,
		Size:           redisinfo.Size,
		InstanceStatus: redisinfo.InstanceTitle,
		RedisShardSize: redisinfo.RedisShardSize,
		NoAuth:         redisinfo.NoAuth,
		PublicIp:       redisinfo.WanAddress,
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}
