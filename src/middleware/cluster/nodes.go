package cluster

import (
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func ClusterConvergeTree(node mysql.ClusterNode) *model.ClusterNodeTables {
	root := &model.ClusterNodeTables{
		CreateTime: node.CreatedAt,
		CluserId:   node.CluserId,
		Flags:      node.Flags,
		Address:    node.Ip + ":" + node.Port,
		LinkState:  node.LinkState,
		RunStatus:  node.RunStatus,
		NodeId:     node.NodeId,
		SlotNumber: node.SlotNumber,
		SlotRange:  node.SlotRange,
	}
	return root
}
func ClusterUpdateTree(mlist []*model.ClusterNodeTables, node mysql.ClusterNode) []*model.ClusterNodeTables {
	var res []*model.ClusterNodeTables
	for _, v := range mlist {
		if v.NodeId == node.MasterId {
			v.Children = append(v.Children, ClusterConvergeTree(node))
			res = append(res, v)
		} else {
			res = append(res, v)
		}
	}
	return res
}
