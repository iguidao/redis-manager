package model

var (
	DefaultName           = make(map[string]string)
	CC                    = "custom_config"                                                                                    // 自行配置的key
	CN                    = "用户自定义默认key"                                                                                       // 自定义key备注note
	TXSECRETID            = "tx_secretid"                                                                                      // 腾讯SECRETID，账号需要开启[QCloudFinanceFullAccess、QcloudRedisFullAccess、QcloudMonitorFullAccess、QcloudDBBRAINFullAccess]权限
	TXSECRETKEY           = "tx_secretkey"                                                                                     // 腾讯SECRETKEY
	TXAPIURL              = "tx_redis_api_url"                                                                                 // 腾讯APIURL
	TXCOSACCESSKEY        = "tx_cos_accesskey"                                                                                 // 腾讯COS的ACCESSKEY
	TXCOSACCESSKEYID      = "tx_cos_accesskeyid"                                                                               // 腾讯COS的ACCESSKEYID
	TXCOSENDPOINTPUB      = "tx_cos_endpointpub"                                                                               // 腾讯COS的ENDPOINTPUB
	ALIAPIURL             = "ali_redis_api_url"                                                                                // 阿里PIURL
	ALIACCESSKEYID        = "ali_accesskeyid"                                                                                  // 阿里accessKeyId
	ALIALIACCESSKEYSECRET = "ali_accesskeysecret"                                                                              // 阿里accessKeySecret
	BGSAVECOMMAND         = "redis_bgsave"                                                                                     // bgsave命令的别名
	CLOUDREFRESH          = "cloud_refresh"                                                                                    // 云redis定时更新时间，使用cron格式
	BOARDCODIS            = "board_codis"                                                                                      // 是否启动自建codis
	BOARDTXREDIS          = "board_txredis"                                                                                    // 是否启动腾讯redis
	BOARDALIREDIS         = "board_aliredis"                                                                                   // 是否启动阿里redis
	BOARDCLUSTER          = "board_cluster"                                                                                    // 是否启动自建redis
	CfgDefault            = [...]string{TXSECRETID, TXSECRETKEY, TXAPIURL, TXCOSACCESSKEY, TXCOSACCESSKEYID, TXCOSENDPOINTPUB} // 默认key列表
)

func init() {
	DefaultName[TXSECRETID] = "腾讯SECRETID"
	DefaultName[TXSECRETKEY] = "腾讯TXSECRETKEY"
	DefaultName[TXAPIURL] = "腾讯REDIS的APIURL"
	DefaultName[TXCOSACCESSKEY] = "腾讯COS的ACCESSKEY"
	DefaultName[TXCOSACCESSKEYID] = "腾讯COS的ACCESSKEYID"
	DefaultName[TXCOSENDPOINTPUB] = "腾讯COS的ENDPOINTPUB"
	DefaultName[BGSAVECOMMAND] = "Redis命令bgsave别名"
	DefaultName[CLOUDREFRESH] = "云redis定时更新时间"
	DefaultName[ALIAPIURL] = "阿里REDIS的APIURL"
	DefaultName[ALIACCESSKEYID] = "阿里accessKeyId"
	DefaultName[ALIALIACCESSKEYSECRET] = "阿里accessKeySecret"
}
