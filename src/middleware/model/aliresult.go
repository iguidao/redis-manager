package model

// Region ali result
type AliRegion struct {
	RequestId string             `json:"RequestId"`
	RegionIds AliRegionRegionIds `json:"RegionIds"`
}
type AliRegionRegionIds struct {
	KVStoreRegion []AliRegionRegionIdsKVStoreRegion `json:"KVStoreRegion"`
}
type AliRegionRegionIdsKVStoreRegion struct {
	RegionId       string `json:"RegionId"`
	RegionEndpoint string `json:"RegionEndpoint"`
	LocalName      string `json:"LocalName"`
}

// redis ali result
type AliRedis struct {
	RequestId  string            `json:"RequestId"`
	TotalCount int               `json:"TotalCount"`
	PageSize   int               `json:"PageSize"`
	PageNumber int               `json:"PageNumber"`
	Instances  AliRedisInstances `json:"RegionIds"`
}
type AliRedisInstances struct {
	KVStoreInstance []AliRedisInstancesKVStoreInstance `json:"KVStoreInstance"`
}
type AliRedisInstancesKVStoreInstance struct {
	Connections         int    `json:"Connections"`
	EndTime             string `json:"EndTime"`
	ResourceGroupId     string `json:"ResourceGroupId"`
	EditionType         string `json:"EditionType"`
	Config              string `json:"Config"`
	Port                int    `json:"Port"`
	GlobalInstanceId    string `json:"GlobalInstanceId"`
	HasRenewChangeOrder string `json:"HasRenewChangeOrder"`
	ConnectionDomain    string `json:"ConnectionDomain"`
	Capacity            int    `json:"Capacity"`
	QPS                 int    `json:"QPS"`
	NetworkType         string `json:"NetworkType"`
	InstanceStatus      string `json:"InstanceStatus"`
	PackageType         string `json:"PackageType"`
	Bandwidth           int    `json:"Bandwidth"`
	InstanceType        string `json:"InstanceType"`
	ArchitectureType    string `json:"ArchitectureType"`
	EngineVersion       string `json:"EngineVersion"`
	UserName            string `json:"UserName"`
	ZoneId              string `json:"ZoneId"`
	InstanceId          string `json:"InstanceId"`
	CreateTime          string `json:"CreateTime"`
	VSwitchId           string `json:"VSwitchId"`
	InstanceClass       string `json:"InstanceClass"`
	IsRds               bool   `json:"IsRds"`
	InstanceName        string `json:"InstanceName"`
	VpcId               string `json:"VpcId"`
	ChargeType          string `json:"ChargeType"`
	NodeType            string `json:"NodeType"`
	RegionId            string `json:"RegionId"`
	ShardCount          int    `json:"ShardCount"`
}
