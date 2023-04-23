package mysql

// get all cluster
func (m *MySQL) GetAllCluster() []ClusterInfo {
	var clusters []ClusterInfo
	m.Find(&clusters)
	return clusters
}

func (m *MySQL) GetClusterNode(cluster string) []ClusterNode {
	var nodes []ClusterNode
	m.Model(nodes).Where("cluser_id = ?", cluster).Find(&nodes)
	return nodes
}

// add cluster
func (m *MySQL) AddCluster(name, nodes, password string) (int, bool) {
	addcluster := &ClusterInfo{
		Name:     "",
		Nodes:    "",
		Password: "",
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
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
		return 0, false
	}
	return addnode.ID, true
	// return gdarticle.ID.String(), true
}
