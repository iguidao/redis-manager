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
