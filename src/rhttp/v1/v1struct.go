package v1

// codis信息的操作
type CodisInfo struct {
	Curl        string `json:"curl"`
	Cname       string `json:"cname"`
	ClusterName string `json:"cluster_name"`
}

// config配置信息
type ConfigInfo struct {
	// Name  string `json:"name"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 操作缓存的指令
type CliQuery struct {
	CacheType   string `json:"cache_type"`
	CacheOp     string `json:"cache_op"`
	ClusterName string `json:"cluster_name"`
	KeyName     string `json:"key_name"`
	CodisUrl    string `json:"codis_url"`
	GroupName   string `json:"group_name"`
	Region      string `json:"region"`
	InstanceId  string `json:"instance_id"`
	ClusterId   string `json:"cluster_id"`
	NodeId      string `json:"node_id"`
}

// 分析大key
type CliRdb struct {
	RdbName  string `json:"rdbname"`
	ServerIp string `json:"serverip"`
}

// old
type AddCluster struct {
	Name     string `json:"name"`
	Nodes    string `json:"nodes"`
	Password string `json:"password"`
}

// user
type UserInfo struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	UserType string `json:"usertype"`
}
type UserPassword struct {
	Old string `json:"old"`
	New string `json:"new"`
}

// casbin rule
type CasbinPolicyJson struct {
	Identity string `json:"identity"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

// cloud
type CloudPassword struct {
	Cloud      string `json:"cloud"`
	Instanceid string `json:"instanceid"`
	Password   string `json:"password"`
}

type TxShardCfg struct {
	Cloud        string `json:"cloud"`
	TxShardType  string `json:"txshardtype"`
	TxShardValue string `json:"txshardvalue"`
}

// cluster nodes table
