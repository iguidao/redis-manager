package v1

type AddCluster struct {
	UserId        int    `json:"user_id"`
	GroupId       int    `json:"group_id"`
	ClusterName   string `json:"cluster_name"`
	RedisPassword string `json:"redis_password"`
	RedisNodes    string `json:"redis_nodes"`
	Environment   string `json:"environment"`
	ClusterNotes  string `json:"cluster_notes"`
}

type CliQuery struct {
	ClusterName string `json:"cluster_name"`
	KeyName     string `json:"key_name"`
	GroupName   string `json:"group_name"`
	Env         string `json:"env"`
}

type CliRdb struct {
	RdbName  string `json:"rdbname"`
	ServerIp string `json:"serverip"`
}
