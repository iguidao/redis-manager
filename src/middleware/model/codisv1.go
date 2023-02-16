package model

// codis的proxy和server操作
type CodisNode struct {
	Curl        string           `json:"curl"`
	ClusterName string           `json:"cluster_name"`
	Proxy       CodisNodeProxy   `json:"proxy"`
	Group       []CodisNodeGroup `json:"group"`
	OpType      string           `json:"op_type"`
}

type CodisNodeProxy struct {
	List []string `json:"list"`
	Port string   `json:"port"`
}

type CodisNodeGroup struct {
	Id   string   `json:"id"`
	List []string `json:"list"`
	Port string   `json:"port"`
}
