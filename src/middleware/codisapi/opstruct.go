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

// /topom?forward= 接口返回数据 =======================================================================================================================================
type Topom struct {
	Version string      `json:"version"`
	Compile string      `json:"compile"`
	Config  TopomConfig `json:"config"`
	Model   TopomModel  `json:"model"`
	Stats   TopomStats  `json:"stats"`
}

type TopomConfig struct {
	Coordinator_name                string `json:"coordinator_name"`
	Coordinator_addr                string `json:"coordinator_addr"`
	Coordinator_auth                string `json:"coordinator_auth"`
	Admin_addr                      string `json:"admin_addr"`
	Product_name                    string `json:"product_name"`
	Migration_method                string `json:"migration_method"`
	Migration_parallel_slots        string `json:"migration_parallel_slots"`
	Migration_async_maxbulks        string `json:"migration_async_maxbulks"`
	Migration_async_maxbytes        string `json:"migration_async_maxbytes"`
	Migration_async_numkeys         string `json:"migration_async_numkeys"`
	Migration_timeout               string `json:"migration_timeout"`
	Sentinel_client_timeout         string `json:"sentinel_client_timeout"`
	Sentinel_quorum                 string `json:"sentinel_quorum"`
	Sentinel_parallel_syncs         string `json:"sentinel_parallel_syncs"`
	Sentinel_down_after             string `json:"sentinel_down_after"`
	Sentinel_failover_timeout       string `json:"sentinel_failover_timeout"`
	Sentinel_notification_script    string `json:"sentinel_notification_script"`
	Sentinel_client_reconfig_script string `json:"sentinel_client_reconfig_script"`
}

type TopomModel struct {
	Token        string `json:"token"`
	Start_time   string `json:"start_time"`
	Admin_addr   string `json:"admin_addr"`
	Product_name string `json:"product_name"`
	Pid          string `json:"pid"`
	Pwd          string `json:"pwd"`
	Sys          string `json:"sys"`
}

type TopomStats struct {
	Closed     bool                 `json:"closed"`
	Slots      []TopomStatsSlots    `json:"slots"`
	Group      TopomStatsGroup      `json:"group"`
	Proxy      TopomStatsProxy      `json:"proxy"`
	SlotAction TopomStatsSlotAction `json:"slot_action"`
	Sentinels  TopomStatsSentinels  `json:"sentinels"`
}

// TopomStats 中 Slots 的参数
type TopomStatsSlots struct {
	Id      int                   `json:"id"`
	GroupId int                   `json:"group_id"`
	Action  TopomStatsSlotsAction `json:"action"`
}
type TopomStatsSlotsAction struct {
	Index    int    `json:"index"`
	State    string `json:"state"`
	TargetId int    `json:"target_id"`
}

// TopomStats 中 Group 的参数
type TopomStatsGroup struct {
	Models []TopomStatsGroupModels         `json:"models"`
	Stats  map[string]TopomStatsGroupStats `json:"stats"`
}
type TopomStatsGroupModels struct {
	Id        int                            `json:"id"`
	Servers   []TopomStatsGroupModelsServers `json:"servers"`
	Promoting TopomStatsGroupModelsPromoting `json:"promoting"`
	OutOfSync bool                           `json:"out_of_sync"`
}
type TopomStatsGroupModelsServers struct {
	Server       string                             `json:"server"`
	Datacenter   string                             `json:"datacenter"`
	Action       TopomStatsGroupModelsServersAction `json:"action"`
	ReplicaGroup bool                               `json:"replica_group"`
}
type TopomStatsGroupModelsServersAction struct {
	State string `json:"state"`
}
type TopomStatsGroupModelsPromoting struct {
}
type TopomStatsGroupStats struct {
	Stats    TopomStatsGroupStatsStats `json:"stats"`
	Unixtime int                       `json:"unixtime"`
}
type TopomStatsGroupStatsStats struct {
	Aof_current_rewrite_time_sec   string `json:"aof_current_rewrite_time_sec"`
	Aof_enabled                    string `json:"aof_enabled"`
	Aof_last_bgrewrite_status      string `json:"aof_last_bgrewrite_status"`
	Aof_last_rewrite_time_sec      string `json:"aof_last_rewrite_time_sec"`
	Aof_last_write_status          string `json:"aof_last_write_status"`
	Aof_rewrite_in_progress        string `json:"aof_rewrite_in_progress"`
	Aof_rewrite_scheduled          string `json:"aof_rewrite_scheduled"`
	Arch_bits                      string `json:"arch_bits"`
	Blocked_clients                string `json:"blocked_clients"`
	Client_biggest_input_buf       string `json:"client_biggest_input_buf"`
	Client_longest_output_list     string `json:"client_longest_output_list"`
	Cluster_enabled                string `json:"cluster_enabled"`
	Config_file                    string `json:"config_file"`
	Connected_clients              string `json:"connected_clients"`
	Connected_slaves               string `json:"connected_slaves"`
	Evicted_keys                   string `json:"evicted_keys"`
	Executable                     string `json:"executable"`
	Expired_keys                   string `json:"expired_keys"`
	Gcc_version                    string `json:"gcc_version"`
	Hz                             string `json:"hz"`
	Instantaneous_input_kbps       string `json:"instantaneous_input_kbps"`
	Instantaneous_ops_per_sec      string `json:"instantaneous_ops_per_sec"`
	Instantaneous_output_kbps      string `json:"instantaneous_output_kbps"`
	Keyspace_hits                  string `json:"keyspace_hits"`
	Keyspace_misses                string `json:"keyspace_misses"`
	Latest_fork_usec               string `json:"latest_fork_usec"`
	Loading                        string `json:"loading"`
	Lru_clock                      string `json:"lru_clock"`
	Master_addr                    string `json:"master_addr"`
	Master_host                    string `json:"master_host"`
	Master_last_io_seconds_ago     string `json:"master_last_io_seconds_ago"`
	Master_link_status             string `json:"master_link_status"`
	Master_port                    string `json:"master_port"`
	Master_repl_offset             string `json:"master_repl_offset"`
	Master_sync_in_progress        string `json:"master_sync_in_progress"`
	Maxmemory                      string `json:"maxmemory"`
	Maxmemory_human                string `json:"maxmemory_human"`
	Maxmemory_policy               string `json:"maxmemory_policy"`
	Mem_allocator                  string `json:"mem_allocator"`
	Mem_fragmentation_ratio        string `json:"mem_fragmentation_ratio"`
	Migrate_cached_sockets         string `json:"migrate_cached_sockets"`
	Multiplexing_api               string `json:"multiplexing_api"`
	Os                             string `json:"os"`
	Process_id                     string `json:"process_id"`
	Pubsub_channels                string `json:"pubsub_channels"`
	Pubsub_patterns                string `json:"pubsub_patterns"`
	Rdb_bgsave_in_progress         string `json:"rdb_bgsave_in_progress"`
	Rdb_changes_since_last_save    string `json:"rdb_changes_since_last_save"`
	Rdb_current_bgsave_time_sec    string `json:"rdb_current_bgsave_time_sec"`
	Rdb_last_bgsave_status         string `json:"rdb_last_bgsave_status"`
	Rdb_last_bgsave_time_sec       string `json:"rdb_last_bgsave_time_sec"`
	Rdb_last_save_time             string `json:"rdb_last_save_time"`
	Redis_build_id                 string `json:"redis_build_id"`
	Redis_git_dirty                string `json:"redis_git_dirty"`
	Redis_git_sha1                 string `json:"redis_git_sha1"`
	Redis_mode                     string `json:"redis_mode"`
	Redis_version                  string `json:"redis_version"`
	Rejected_connections           string `json:"rejected_connections"`
	Repl_backlog_active            string `json:"repl_backlog_active"`
	Repl_backlog_first_byte_offset string `json:"repl_backlog_first_byte_offset"`
	Repl_backlog_histlen           string `json:"repl_backlog_histlen"`
	Repl_backlog_size              string `json:"repl_backlog_size"`
	Role                           string `json:"role"`
	Run_id                         string `json:"run_id"`
	Slave_priority                 string `json:"slave_priority"`
	Slave_read_only                string `json:"slave_read_only"`
	Slave_repl_offset              string `json:"slave_repl_offset"`
	Sync_full                      string `json:"sync_full"`
	Sync_partial_err               string `json:"sync_partial_err"`
	Sync_partial_ok                string `json:"sync_partial_ok"`
	Tcp_port                       string `json:"tcp_port"`
	Total_commands_processed       string `json:"total_commands_processed"`
	Total_connections_received     string `json:"total_connections_received"`
	Total_net_input_bytes          string `json:"total_net_input_bytes"`
	Total_net_output_bytes         string `json:"total_net_output_bytes"`
	Total_system_memory            string `json:"total_system_memory"`
	Total_system_memory_human      string `json:"total_system_memory_human"`
	Uptime_in_days                 string `json:"uptime_in_days"`
	Uptime_in_seconds              string `json:"uptime_in_seconds"`
	Used_cpu_sys                   string `json:"used_cpu_sys"`
	Used_cpu_sys_children          string `json:"used_cpu_sys_children"`
	Used_cpu_user                  string `json:"used_cpu_user"`
	Used_cpu_user_children         string `json:"used_cpu_user_children"`
	Used_memory                    string `json:"used_memory"`
	Used_memory_human              string `json:"used_memory_human"`
	Used_memory_lua                string `json:"used_memory_lua"`
	Used_memory_lua_human          string `json:"used_memory_lua_human"`
	Used_memory_peak               string `json:"used_memory_peak"`
	Used_memory_peak_human         string `json:"used_memory_peak_human"`
	Used_memory_rss                string `json:"used_memory_rss"`
	Used_memory_rss_human          string `json:"used_memory_rss_human"`
}

// TopomStats 中 Proxy 的参数
type TopomStatsProxy struct {
	Models []TopomStatsProxyModels         `json:"models"`
	Stats  map[string]TopomStatsProxyStats `json:"stats"`
}
type TopomStatsProxyModels struct {
	Id          int    `json:"id"`
	Token       string `json:"token"`
	StartTime   string `json:"start_time"`
	AdminAddr   string `json:"admin_addr"`
	ProtoType   string `json:"proto_type"`
	ProxyAddr   string `json:"proxy_addr"`
	JodisPath   string `json:"jodis_path"`
	ProductName string `json:"product_name"`
	Pid         int    `json:"pid"`
	Pwd         string `json:"pwd"`
	Sys         string `json:"sys"`
	Hostname    string `json:"hostname"`
	Datacenter  string `json:"datacenter"`
}
type TopomStatsProxyStats struct {
	Stats    TopomStatsProxyStatsStats `json:"stats"`
	Unixtime int                       `json:"unixtime"`
}
type TopomStatsProxyStatsStats struct {
	Online    bool                               `json:"online"`
	Closed    bool                               `json:"closed"`
	Sentinels TopomStatsProxyStatsStatsSentinels `json:"sentinels"`
	Ops       TopomStatsProxyStatsStatsOps       `json:"ops"`
	Sessions  TopomStatsProxyStatsStatsSessions  `json:"sessions"`
	Rusage    TopomStatsProxyStatsStatsRusage    `json:"rusage"`
	Backend   TopomStatsProxyStatsStatsBackend   `json:"backend"`
}
type TopomStatsProxyStatsStatsSentinels struct {
	Servers []string          `json:"servers"`
	Masters map[string]string `json:"masters"`
}
type TopomStatsProxyStatsStatsOps struct {
	Total int                             `json:"total"`
	Fails int                             `json:"fails"`
	Redis TopomStatsProxyStatsStatsOpsqps `json:"redis"`
	Qps   int                             `json:"qps"`
}
type TopomStatsProxyStatsStatsOpsqps struct {
	Errors int `json:"errors"`
}
type TopomStatsProxyStatsStatsSessions struct {
	Total int `json:"total"`
	Alive int `json:"alive"`
}
type TopomStatsProxyStatsStatsRusage struct {
	Now string                             `json:"now"`
	Cpu int                                `json:"cpu"`
	Mem int                                `json:"mem"`
	Raw TopomStatsProxyStatsStatsRusageRaw `json:"raw"`
}
type TopomStatsProxyStatsStatsRusageRaw struct {
	Utime       int `json:"utime"`
	Stime       int `json:"stime"`
	Cutime      int `json:"cutime"`
	Cstime      int `json:"cstime"`
	Num_threads int `json:"num_threads"`
	Vm_size     int `json:"vm_size"`
	Vm_rss      int `json:"vm_rss"`
}
type TopomStatsProxyStatsStatsBackend struct {
	Primary_only bool `json:"primary_only"`
}

// TopomStats 中 SlotAction 的参数
type TopomStatsSlotAction struct {
	Interval int                          `json:"interval"`
	Disabled bool                         `json:"disabled"`
	Progress TopomStatsSlotActionProgress `json:"progress"`
	Executor int                          `json:"executor"`
}
type TopomStatsSlotActionProgress struct {
	Status string `json:"status"`
}

// TopomStats 中 Sentinels 的参数
type TopomStatsSentinels struct {
	Model   TopomStatsSentinelsModel            `json:"model"`
	Stats   map[string]TopomStatsSentinelsStats `json:"stats"`
	Masters map[string]string                   `json:"masters"`
}
type TopomStatsSentinelsModel struct {
	Servers   []string `json:"servers"`
	OutOfSync bool     `json:"out_of_sync"`
}
type TopomStatsSentinelsStats struct {
	Stats    TopomStatsSentinelsStatsStats               `json:"stats"`
	Sentinel map[string]TopomStatsSentinelsStatsSentinel `json:"sentinel"`
	Unixtime int                                         `json:"unixtime"`
}
type TopomStatsSentinelsStatsStats struct {
	Arch_bits                       string `json:"arch_bits"`
	Blocked_clients                 string `json:"blocked_clients"`
	Client_biggest_input_buf        string `json:"client_biggest_input_buf"`
	Client_longest_output_list      string `json:"client_longest_output_list"`
	Config_file                     string `json:"config_file"`
	Connected_clients               string `json:"connected_clients"`
	Evicted_keys                    string `json:"evicted_keys"`
	Executable                      string `json:"executable"`
	Expired_keys                    string `json:"expired_keys"`
	Gcc_version                     string `json:"gcc_version"`
	Hz                              string `json:"hz"`
	Instantaneous_input_kbps        string `json:"instantaneous_input_kbps"`
	Instantaneous_ops_per_sec       string `json:"instantaneous_ops_per_sec"`
	Instantaneous_output_kbps       string `json:"instantaneous_output_kbps"`
	Keyspace_hits                   string `json:"keyspace_hits"`
	Keyspace_misses                 string `json:"keyspace_misses"`
	Latest_fork_usec                string `json:"latest_fork_usec"`
	Lru_clock                       string `json:"lru_clock"`
	Master0                         string `json:"master0"`
	Master1                         string `json:"master1"`
	Master2                         string `json:"master2"`
	Master3                         string `json:"master3"`
	Master4                         string `json:"master4"`
	Migrate_cached_sockets          string `json:"migrate_cached_sockets"`
	Multiplexing_api                string `json:"multiplexing_api"`
	Os                              string `json:"os"`
	Process_id                      string `json:"process_id"`
	Pubsub_channels                 string `json:"pubsub_channels"`
	Pubsub_patterns                 string `json:"pubsub_patterns"`
	Redis_build_id                  string `json:"redis_build_id"`
	Redis_git_dirty                 string `json:"redis_git_dirty"`
	Redis_git_sha1                  string `json:"redis_git_sha1"`
	Redis_mode                      string `json:"redis_mode"`
	Redis_version                   string `json:"redis_version"`
	Rejected_connections            string `json:"rejected_connections"`
	Run_id                          string `json:"run_id"`
	Sentinel_masters                string `json:"sentinel_masters"`
	Sentinel_running_scripts        string `json:"sentinel_running_scripts"`
	Sentinel_scripts_queue_length   string `json:"sentinel_scripts_queue_length"`
	Sentinel_simulate_failure_flags string `json:"sentinel_simulate_failure_flags"`
	Sentinel_tilt                   string `json:"sentinel_tilt"`
	Sync_full                       string `json:"sync_full"`
	Sync_partial_err                string `json:"sync_partial_err"`
	Sync_partial_ok                 string `json:"sync_partial_ok"`
	Tcp_port                        string `json:"tcp_port"`
	Total_commands_processed        string `json:"total_commands_processed"`
	Total_connections_received      string `json:"total_connections_received"`
	Total_net_input_bytes           string `json:"total_net_input_bytes"`
	Total_net_output_bytes          string `json:"total_net_output_bytes"`
	Uptime_in_days                  string `json:"uptime_in_days"`
	Uptime_in_seconds               string `json:"uptime_in_seconds"`
	Used_cpu_sys                    string `json:"used_cpu_sys"`
	Used_cpu_sys_children           string `json:"used_cpu_sys_children"`
	Used_cpu_user                   string `json:"used_cpu_user"`
	Used_cpu_user_children          string `json:"used_cpu_user_children"`
}
type TopomStatsSentinelsStatsSentinel struct {
	Master TopomStatsSentinelsStatsSentinelMaster   `json:"master"`
	Slaves []TopomStatsSentinelsStatsSentinelSlaves `json:"slaves"`
}
type TopomStatsSentinelsStatsSentinelMaster struct {
	Configepoch           string `json:"config-epoch"`
	Downaftermilliseconds string `json:"down-after-milliseconds"`
	Failovertimeout       string `json:"failover-timeout"`
	Flags                 string `json:"flags"`
	Inforefresh           string `json:"info-refresh"`
	Ip                    string `json:"ip"`
	Lastokpingreply       string `json:"last-ok-ping-reply"`
	Lastpingreply         string `json:"last-ping-reply"`
	Lastpingsent          string `json:"last-ping-sent"`
	Linkpendingcommands   string `json:"link-pending-commands"`
	Linkrefcount          string `json:"link-refcount"`
	Name                  string `json:"name"`
	Numothersentinels     string `json:"num-other-sentinels"`
	Numslaves             string `json:"num-slaves"`
	Parallelsyncs         string `json:"parallel-syncs"`
	Port                  string `json:"port"`
	Quorum                string `json:"quorum"`
	Rolereported          string `json:"role-reported"`
	Rolereportedtime      string `json:"role-reported-time"`
	Runid                 string `json:"runid"`
}
type TopomStatsSentinelsStatsSentinelSlaves struct {
	Downaftermilliseconds string `json:"down-after-milliseconds"`
	Flags                 string `json:"flags"`
	Inforefresh           string `json:"info-refresh"`
	Ip                    string `json:"ip"`
	Lastokpingreply       string `json:"last-ok-ping-reply"`
	Lastpingreply         string `json:"last-ping-reply"`
	Lastpingsent          string `json:"last-ping-sent"`
	Linkpendingcommands   string `json:"link-pending-commands"`
	Linkrefcount          string `json:"link-refcount"`
	Masterhost            string `json:"master-host"`
	Masterlinkdowntime    string `json:"master-link-down-time"`
	Masterlinktatus       string `json:"master-link-status"`
	Masterport            string `json:"master-port"`
	Name                  string `json:"name"`
	Port                  string `json:"port"`
	Rolereported          string `json:"role-reported"`
	Rolereportedtime      string `json:"role-reported-time"`
	Runid                 string `json:"runid"`
	Slavepriority         string `json:"slave-priority"`
	Slavereploffset       string `json:"slave-repl-offset"`
}
