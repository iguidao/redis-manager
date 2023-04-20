package hsc

// 提供错误信息
var MsgFlags = map[int]string{
	SUCCESS:                        "succes",
	ERROR:                          "error",
	NOT_FOUND:                      "Not Found",
	INVALID_PARAMS:                 "invalid params",
	NOT_PROMISE:                    "no promise",
	Method_FAILS:                   "Method Fails",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Check Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Create Fail",
	ERROR_AUTH:                     "认证失败",

	ERROR_NO_CONNEC:       "链接目标Redis异常",
	WARN_NO_USE:           "功能会造成慢查询，暂时下线",
	WARN_CLICK_REPEATEDLY: "兄弟，你点的太快了，上一个还没结束，等一下哈！",
	WARN_BACKGROUND:       "没有找到这些数据，后台已经再加载，请稍后重试一下。",
	ERROR_BACKGROUND:      "后台加载失败，请联系管理员",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
