package hsc

// 提供错误信息
var MsgFlags = map[int]string{
	SUCCESS:        "succes",
	ERROR:          "error",
	INVALID_PARAMS: "invalid params",
	NOT_PROMISE:    "no promise",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token Check Fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token Timeout",
	ERROR_AUTH_TOKEN:               "Token Create Fail",
	ERROR_AUTH:                     "Token Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
