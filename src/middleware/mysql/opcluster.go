package mysql

// get all cluster
func (m *MySQL) GetAllCluster() []ClusterInfo {
	var clusters []ClusterInfo
	m.Find(&clusters)
	return clusters
}

// func (m *MySQL) GetCluster(page int, size int, status int) (clusterinfo []ClusterInfo) {
// 	// m.Where(maps).Offset(page).Limit(size).Find(&clusterinfo)
// 	m.Where("state = ?", status).Offset(page).Limit(size).Find(&clusterinfo)
// 	return
// }

// add cluster
func (m *MySQL) AddCluster(ArticleTitle string, ArticleContent string, AuthorId string) (uint, bool) {
	addcluster := &ClusterInfo{
		ClusterName:    "",
		ClusterMode:    "", // 集群(Cluster)；单点(Single)；哨兵(Sentinel)
		ClusterVersion: "",
		NodesAll:       0,
		NodesMaster:    0,
		NodesSlave:     0,
		RedisPassword:  "",
		Environment:    "", // 主机 Machine；容器 Container
		From:           "", //导入Import；平台创建Self
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}
