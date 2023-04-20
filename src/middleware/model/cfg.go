package model

var (
	DefaultName      = make(map[string]string)
	CC               = "custom_config"                                                                                    // 自行配置的key
	CN               = "用户自定义默认key"                                                                                       // 自定义key备注note
	TXSECRETID       = "tx_secretid"                                                                                      // 腾讯SECRETID，账号需要开启[QCloudFinanceFullAccess、QcloudRedisFullAccess、QcloudMonitorFullAccess]权限
	TXSECRETKEY      = "tx_secretkey"                                                                                     // 腾讯SECRETKEY
	TXAPIURL         = "tx_redis_api_url"                                                                                 // 腾讯APIURL
	TXCOSACCESSKEY   = "tx_cos_accesskey"                                                                                 // 腾讯COS的ACCESSKEY
	TXCOSACCESSKEYID = "tx_cos_accesskeyid"                                                                               // 腾讯COS的ACCESSKEYID
	TXCOSENDPOINTPUB = "tx_cos_endpointpub"                                                                               // 腾讯COS的ENDPOINTPUB
	BGSAVECOMMAND    = "redis_bgsave"                                                                                     //bgsave命令的别名
	CfgDefault       = [...]string{TXSECRETID, TXSECRETKEY, TXAPIURL, TXCOSACCESSKEY, TXCOSACCESSKEYID, TXCOSENDPOINTPUB} // 默认key列表
)

func init() {
	DefaultName[TXSECRETID] = "腾讯SECRETID"
	DefaultName[TXSECRETKEY] = "腾讯TXSECRETKEY"
	DefaultName[TXAPIURL] = "腾讯REDIS的APIURL"
	DefaultName[TXCOSACCESSKEY] = "腾讯COS的ACCESSKEY"
	DefaultName[TXCOSACCESSKEYID] = "腾讯COS的ACCESSKEYID"
	DefaultName[TXCOSENDPOINTPUB] = "腾讯COS的ENDPOINTPUB"
	DefaultName[BGSAVECOMMAND] = "Redis命令bgsave别名"
}
