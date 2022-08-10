package codisapi

type ResultForward struct {
	Version string        `json:"version"`
	Compile string        `json:"compile"`
	Config  ForwardConfig `json:"config"`
	Model   ForwardModel  `json:"model"`
	Stats   ForwardStats  `json:"stats"`
}

type ForwardConfig struct {
}

type ForwardModel struct {
}

type ForwardStats struct {
	Closed    bool           `json:"closed"`
	Group     StatsGroup     `json:"group"`
	Proxy     StatsProxy     `json:"proxy"`
	Sentinels StatsSentinels `json:"sentinels"`
}

type StatsGroup struct {
	Models []GroupModels `json:"models"`
}

type GroupModels struct {
	Id      int             `json:"id"`
	Servers []ModelsServers `json:"servers"`
}

type ModelsServers struct {
	Server string        `json:"server"`
	Action ServersAction `json:"action"`
}

type ServersAction struct {
	State string `json:"state"`
}
type StatsProxy struct {
	Models []ProxyModels `json:"models"`
}
type ProxyModels struct {
	ProxyAddr string `json:"proxy_addr"`
}

type StatsSentinels struct {
	Masters interface{} `json:"masters"`
}
