package model

// tx redis list result
type TxL struct {
	Response TxLResponse `json:"Response"`
}

type TxLResponse struct {
	TotalCount  int                      `json:"TotalCount"`
	InstanceSet []TxLResponseInstanceSet `json:"InstanceSet"`
	RequestId   string                   `json:"RequestId"`
}

type TxLResponseInstanceSet struct {
	InstanceName            string                          `json:"InstanceName"`
	InstanceId              string                          `json:"InstanceId"`
	Appid                   int                             `json:"Appid"`
	ProjectId               int                             `json:"ProjectId"`
	RegionId                int                             `json:"RegionId"`
	ZoneId                  int                             `json:"ZoneId"`
	VpcId                   int                             `json:"VpcId"`
	SubnetId                int                             `json:"SubnetId"`
	Status                  int                             `json:"Status"`
	WanIp                   string                          `json:"WanIp"`
	Port                    int                             `json:"Port"`
	Createtime              string                          `json:"Createtime"`
	Size                    int                             `json:"Size"`
	SizeUsed                int                             `json:"SizeUsed"`
	Type                    int                             `json:"Type"`
	AutoRenewFlag           int                             `json:"AutoRenewFlag"`
	DeadlineTime            string                          `json:"DeadlineTime"`
	Engine                  string                          `json:"Engine"`
	ProductType             string                          `json:"ProductType"`
	UniqVpcId               string                          `json:"UniqVpcId"`
	UniqSubnetId            string                          `json:"UniqSubnetId"`
	BillingMode             int                             `json:"BillingMode"`
	InstanceTitle           string                          `json:"InstanceTitle"`
	OfflineTime             string                          `json:"OfflineTime"`
	SubStatus               int                             `json:"SubStatus"`
	RedisShardSize          int                             `json:"RedisShardSize"`
	RedisShardNum           int                             `json:"RedisShardNum"`
	RedisReplicasNum        int                             `json:"RedisReplicasNum"`
	PriceId                 int                             `json:"PriceId"`
	CloseTime               string                          `json:"CloseTime"`
	SlaveReadWeight         int                             `json:"SlaveReadWeight"`
	ProjectName             string                          `json:"ProjectName"`
	NoAuth                  bool                            `json:"NoAuth"`
	ClientLimit             int                             `json:"ClientLimit"`
	DtsStatus               int                             `json:"DtsStatus"`
	NetLimit                int                             `json:"NetLimit"`
	PasswordFree            int                             `json:"PasswordFree"`
	Vip6                    string                          `json:"Vip6"`
	ReadOnly                int                             `json:"ReadOnly"`
	RemainBandwidthDuration string                          `json:"RemainBandwidthDuration"`
	DiskSize                int                             `json:"DiskSize"`
	MonitorVersion          string                          `json:"MonitorVersion"`
	ClientLimitMin          int                             `json:"ClientLimitMin"`
	ClientLimitMax          int                             `json:"ClientLimitMax"`
	NodeSet                 []TxLResponseInstanceSetNodeSet `json:"NodeSet"`
	Region                  string                          `json:"Region"`
	WanAddress              string                          `json:"WanAddress"`
	PolarisServer           string                          `json:"PolarisServer"`
	CurrentProxyVersion     string                          `json:"CurrentProxyVersion"`
	CurrentRedisVersion     string                          `json:"CurrentRedisVersion"`
	UpgradeProxyVersion     string                          `json:"UpgradeProxyVersion"`
	UpgradeRedisVersion     string                          `json:"UpgradeRedisVersion"`
}

type TxLResponseInstanceSetNodeSet struct {
	NodeType int    `json:"NodeType"`
	NodeId   int    `json:"NodeId"`
	ZoneId   int    `json:"ZoneId"`
	ZoneName string `json:"ZoneName"`
}

// Region tx result
type TxRegion struct {
	Response TxRegionResponse `json:"Response"`
}

type TxRegionResponse struct {
	TotalCount int                         `json:"TotalCount"`
	RegionSet  []TxRegionResponseRegionSet `json:"RegionSet"`
	RequestId  string                      `json:"RequestId"`
}
type TxRegionResponseRegionSet struct {
	Region      string `json:"Region"`
	RegionName  string `json:"RegionName"`
	RegionState string `json:"RegionState"`
}

// Hot key resultn
type TxHotKey struct {
	Response TxHotKeyResponse `json:"Response"`
}

type TxHotKeyResponse struct {
	Data      []TxHotKeyResponseData `json:"Data"`
	RequestId string                 `json:"RequestId"`
}
type TxHotKeyResponseData struct {
	Count int    `json:"Count"`
	Key   string `json:"key"`
	Type  string `json:"Type"`
}

// Slow Key result
type TxProxySlowKey struct {
	Response TxProxySlowKeyResponse `json:"Response"`
}
type TxProxySlowKeyResponse struct {
	InstanceProxySlowLogDetail []TxProxySlowKeyResponseInstanceProxySlowLogDetail `json:"InstanceProxySlowLogDetail"`
	RequestId                  string                                             `json:"RequestId"`
	TotalCount                 int                                                `json:"TotalCount"`
}
type TxProxySlowKeyResponseInstanceProxySlowLogDetail struct {
	Client      string `json:"Client"`
	Command     string `json:"Command"`
	CommandLine string `json:"CommandLine"`
	Duration    int    `json:"Duration"`
	ExecuteTime string `json:"ExecuteTime"`
}
type TxRedisSlowKey struct {
	Response TxRedisSlowKeyResponse `json:"Response"`
}

type TxRedisSlowKeyResponse struct {
	InstanceSlowlogDetail []TxRedisSlowKeyResponseInstanceSlowlogDetail `json:"InstanceSlowlogDetail"`
	RequestId             string                                        `json:"RequestId"`
	TotalCount            int                                           `json:"TotalCount"`
}
type TxRedisSlowKeyResponseInstanceSlowlogDetail struct {
	Client      string `json:"Client"`
	Command     string `json:"Command"`
	CommandLine string `json:"CommandLine"`
	Duration    int    `json:"Duration"`
	ExecuteTime string `json:"ExecuteTime"`
	Node        string `json:"Node"`
}
