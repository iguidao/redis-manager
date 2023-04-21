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
}

// 分析大key
type CliRdb struct {
	RdbName  string `json:"rdbname"`
	ServerIp string `json:"serverip"`
}

// old
type AddCluster struct {
	UserId        int    `json:"user_id"`
	GroupId       int    `json:"group_id"`
	ClusterName   string `json:"cluster_name"`
	RedisPassword string `json:"redis_password"`
	RedisNodes    string `json:"redis_nodes"`
	Environment   string `json:"environment"`
	ClusterNotes  string `json:"cluster_notes"`
}

// user
type UserInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	UserType string `json:"usertype"`
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
