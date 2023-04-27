package model

// codis的proxy和server操作
type CodisChangeNode struct {
	Curl        string   `json:"curl"`
	ClusterName string   `json:"cluster_name"`
	AddProxy    []string `json:"add_proxy"`
	AddServer   []string `json:"add_server"`
	DelProxy    int      `json:"del_proxy"`
	DelGroup    int      `json:"del_group"`
	OpType      string   `json:"op_type"`
}
