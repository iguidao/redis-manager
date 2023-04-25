package mysql

import "github.com/iguidao/redis-manager/src/middleware/logger"

// get all cluster
func (m *MySQL) GetAllCluster() []ClusterInfo {
	var clusters []ClusterInfo
	m.Find(&clusters)
	return clusters
}
func (m *MySQL) GetClusterNumber() int64 {
	var clusters []ClusterInfo
	var count int64
	m.Model(clusters).Find(&clusters).Count(&count)
	return count
}
func (m *MySQL) GetClusterAddress(id string) (string, string) {
	var clusterinfo *ClusterInfo
	m.Where("id = ?", id).First(&clusterinfo)
	return clusterinfo.Nodes, clusterinfo.Password
}
func (m *MySQL) GetClusterPassword(id string) string {
	var clusterinfo *ClusterInfo
	m.Where("id = ?", id).First(&clusterinfo)
	return clusterinfo.Password
}

// add cluster
func (m *MySQL) AddCluster(name, nodes, password string) (int, bool) {
	addcluster := &ClusterInfo{
		Name:     name,
		Nodes:    nodes,
		Password: password,
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		logger.Error("Mysql add cluster error:", result.Error)
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}

func (m *MySQL) GetClusterNode(cluster string) []ClusterNode {
	var nodes []ClusterNode
	m.Model(nodes).Where("cluser_id = ?", cluster).Find(&nodes)
	return nodes
}
func (m *MySQL) GetClusterNodeMaster(cluster string) []ClusterNode {
	var nodes []ClusterNode
	m.Model(nodes).Where("cluser_id = ? AND flags = ?", cluster, "master").Find(&nodes)
	return nodes
}

func (m *MySQL) GetClusterNodeMasterAddress(nodeid string) string {
	var nodes *ClusterNode
	m.Where("node_id = ?", nodeid).First(&nodes)
	return nodes.Ip + ":" + nodes.Port
}
func (m *MySQL) GetClusterNodeSlaverAddress(nodeid string) string {
	var nodes *ClusterNode
	m.Where("master_id = ?", nodeid).First(&nodes)
	return nodes.Ip + ":" + nodes.Port
}
func (m *MySQL) AddClusterNode(nodeid, ip, port, flags, masterid, linkstate, slotrange string, clusterid, slotnumber int) (int, bool) {
	addnode := &ClusterNode{
		CluserId:   clusterid,
		NodeId:     nodeid,
		Ip:         ip,
		Port:       port,
		Flags:      flags,
		MasterId:   masterid,
		LinkState:  linkstate,
		SlotRange:  slotrange,
		SlotNumber: slotnumber,
	}
	result := m.Create(&addnode)
	if result.Error != nil {
		logger.Error("Mysql add cluster node error:", result.Error)
		return 0, false
	}
	return addnode.ID, true
	// return gdarticle.ID.String(), true
}
