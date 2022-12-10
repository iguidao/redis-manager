package hsc

// 提供错误信息
var MsgFlags = map[int]string{
	SUCCESS:        "succes",
	ERROR:          "error",
	NOT_FOUND:      "Not Found",
	INVALID_PARAMS: "invalid params",
	NOT_PROMISE:    "no promise",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Check Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Create Fail",
	ERROR_AUTH:                     "Token Error",

	NO_CONNECT_CODIS: "Connect Codis Fail",
	NO_USE:           "功能会造成慢查询，暂时下线",
	CLICK_REPEATEDLY: "兄弟，你点的太快了，上一个还没结束，等一下哈！",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
