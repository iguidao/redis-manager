package hsc

// 提供错误信息
var MsgFlags = map[int]string{
	SUCCESS:                        "succes",
	ERROR:                          "error",
	NOT_FOUND:                      "Not Found",
	INVALID_PARAMS:                 "invalid params",
	NOT_PROMISE:                    "鉴权失败",
	Method_FAILS:                   "Method Fails",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Check Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Create Fail",
	ERROR_AUTH:                     "认证失败",

	ERROR_NO_CONNEC:       "链接目标Redis异常",
	ERROR_CLOUD_CONNECT:   "链接云服务异常",
	ERROR_CLOUD_GET:       "获取云资源异常",
	ERROR_WRITE_MYSQL:     "数据库操作异常",
	ERROR_BACKGROUND:      "后台加载异常，请联系管理员",
	WARN_NO_USE:           "功能会造成慢查询，暂时下线",
	WARN_CLICK_REPEATEDLY: "兄弟，你点的太快了，上一个还没结束，等一下哈！",
	WARN_BACKGROUND:       "没有找到这些数据，后台已经再加载，请稍后重试一下。",
	WARN_NOT_FOUND_CLOUD:  "暂时不支持这个云操作",
	WARN_NOT_PROMISE_RULE: "没有权限，请找管理员开放",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
